package basic

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
	"sync"
	"time"

	"github.com/percona/rds_exporter/config"
)

var (
	Period = 60 * time.Second
	Delay  = 600 * time.Second
	Range  = 600 * time.Second
)

type Scraper struct {
	// params
	instance  *config.Instance
	collector *Collector
	ch        chan<- prometheus.Metric

	// internal
	svc         *cloudwatch.CloudWatch
	constLabels prometheus.Labels
}

func NewScraper(instance *config.Instance, collector *Collector, ch chan<- prometheus.Metric) *Scraper {
	// Create CloudWatch client
	sess, _ := collector.sessions.GetSession(instance.Region, instance.Instance)
	if sess == nil {
		return nil
	}
	svc := cloudwatch.New(sess)

	constLabels := prometheus.Labels{
		"region":   instance.Region,
		"instance": instance.Instance,
	}
	for n, v := range instance.Labels {
		if v == "" {
			delete(constLabels, n)
		} else {
			constLabels[n] = v
		}
	}

	return &Scraper{
		// params
		instance:  instance,
		collector: collector,
		ch:        ch,

		// internal
		svc:         svc,
		constLabels: constLabels,
	}
}

func getLatestDatapoint(datapoints []*cloudwatch.Datapoint) *cloudwatch.Datapoint {
	var latest *cloudwatch.Datapoint = nil

	for dp := range datapoints {
		if latest == nil || latest.Timestamp.Before(*datapoints[dp].Timestamp) {
			latest = datapoints[dp]
		}
	}

	return latest
}

// Scrape makes the required calls to AWS CloudWatch by using the parameters in the Collector.
// Once converted into Prometheus format, the metrics are pushed on the ch channel.
func (s *Scraper) Scrape() {
	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(len(s.collector.metrics))
	for _, metric := range s.collector.metrics {
		metric := metric
		go func() {
			defer wg.Done()

			if err := s.scrapeMetric(metric); err != nil {
				level.Error(s.collector.l).Log("metric", metric.cwName, "error", err)
			}
		}()
	}
}

func (s *Scraper) scrapeMetric(metric Metric) error {
	now := time.Now()
	end := now.Add(-Delay)

	// If metric.statistics is empty, default to all
	stats := metric.statistics
	if stats == nil || len(stats) == 0 {
		stats = []string{"Average", "Sum", "Minimum", "Maximum"}
	}

	params := &cloudwatch.GetMetricStatisticsInput{
		EndTime:    aws.Time(end),
		StartTime:  aws.Time(end.Add(-Range)),
		Period:     aws.Int64(int64(Period.Seconds())),
		MetricName: aws.String(metric.cwName),
		Namespace:  aws.String("AWS/RDS"),
		Dimensions: []*cloudwatch.Dimension{
			{
				Name:  aws.String("DBInstanceIdentifier"),
				Value: aws.String(s.instance.Instance),
			},
		},
		Statistics: aws.StringSlice(stats),
	}

	resp, err := s.svc.GetMetricStatistics(params)
	if err != nil {
		return err
	}

	if len(resp.Datapoints) == 0 {
		return nil
	}

	dp := getLatestDatapoint(resp.Datapoints)
	if dp == nil {
		return nil
	}

	// For each requested statistic, build and send the Prometheus metric
	for _, stat := range stats {
		var value float64

		switch stat {
		case "Average":
			value = aws.Float64Value(dp.Average)
		case "Sum":
			value = aws.Float64Value(dp.Sum)
		case "Maximum":
			value = aws.Float64Value(dp.Maximum)
		case "Minimum":
			value = aws.Float64Value(dp.Minimum)
		default:
			continue
		}

		switch metric.cwName {
		case "EngineUptime":
			value = float64(time.Now().Unix() - int64(value))
		}

		// Append the statistic name to help identify them in Prometheus.
		lowerStat := strings.ToLower(stat)
		nameWithStat := metric.prometheusName + "_" + lowerStat
		helpWithStat := metric.prometheusHelp + " (" + lowerStat + ")"

		// Emit the Prometheus metric
		s.ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(nameWithStat, helpWithStat, nil, s.constLabels),
			prometheus.GaugeValue,
			value,
		)
	}

	return nil
}
