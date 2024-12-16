package basic

var Metrics = []Metric{
	{
		cwName:         "ActiveTransactions",
		prometheusName: "aws_rds_active_transactions",
		prometheusHelp: "ActiveTransactions",
	},
	{
		cwName:         "AuroraBinlogReplicaLag",
		prometheusName: "aws_rds_aurora_binlog_replica_lag",
		prometheusHelp: "AuroraBinlogReplicaLag",
	},
	{
		cwName:         "AuroraReplicaLag",
		prometheusName: "aws_rds_aurora_replica_lag",
		prometheusHelp: "AuroraReplicaLag",
	},
	{
		cwName:         "AuroraReplicaLagMaximum",
		prometheusName: "aws_rds_aurora_replica_lag_maximum",
		prometheusHelp: "AuroraReplicaLagMaximum",
	},
	{
		cwName:         "AuroraReplicaLagMinimum",
		prometheusName: "aws_rds_aurora_replica_lag_minimum",
		prometheusHelp: "AuroraReplicaLagMinimum",
	},
	{
		cwName:         "BinLogDiskUsage",
		prometheusName: "aws_rds_bin_log_disk_usage",
		prometheusHelp: "The amount of disk space occupied by binary logs on the master. Applies to MySQL read replicas. Units: Bytes",
	},
	{
		cwName:         "BlockedTransactions",
		prometheusName: "aws_rds_blocked_transactions",
		prometheusHelp: "BlockedTransactions",
	},
	{
		cwName:         "BufferCacheHitRatio",
		prometheusName: "aws_rds_buffer_cache_hit_ratio",
		prometheusHelp: "BufferCacheHitRatio",
	},
	{
		cwName:         "BurstBalance",
		prometheusName: "aws_rds_burst_balance",
		prometheusHelp: "The percent of General Purpose SSD (gp2) burst-bucket I/O credits available. Units: Percent",
	},
	{
		cwName:         "CPUCreditBalance",
		prometheusName: "aws_rds_cpu_credit_balance",
		prometheusHelp: "[T2 instances] The number of CPU credits available for the instance to burst beyond its base CPU utilization. Credits are stored in the credit balance after they are earned and removed from the credit balance after they expire. Credits expire 24 hours after they are earned. CPU credit metrics are available only at a 5 minute frequency. Units: Count",
	},
	{
		cwName:         "CPUCreditUsage",
		prometheusName: "aws_rds_cpu_credit_usage",
		prometheusHelp: "[T2 instances] The number of CPU credits consumed by the instance. One CPU credit equals one vCPU running at 100% utilization for one minute or an equivalent combination of vCPUs, utilization, and time (for example, one vCPU running at 50% utilization for two minutes or two vCPUs running at 25% utilization for two minutes). CPU credit metrics are available only at a 5 minute frequency. If you specify a period greater than five minutes, use the Sum statistic instead of the Average statistic. Units: Count",
	},
	{
		cwName:         "CPUUtilization",
		prometheusName: "node_cpu",
		prometheusHelp: "The percentage of CPU utilization. Units: Percent",
	},
	{
		cwName:         "CommitLatency",
		prometheusName: "aws_rds_commit_latency",
		prometheusHelp: "CommitLatency",
	},
	{
		cwName:         "CommitThroughput",
		prometheusName: "aws_rds_commit_throughput",
		prometheusHelp: "CommitThroughput",
	},
	{
		cwName:         "DDLLatency",
		prometheusName: "aws_rds_ddl_latency",
		prometheusHelp: "DDLLatency",
	},
	{
		cwName:         "DDLThroughput",
		prometheusName: "aws_rds_ddl_throughput",
		prometheusHelp: "DDLThroughput",
	},
	{
		cwName:         "DMLLatency",
		prometheusName: "aws_rds_dml_latency",
		prometheusHelp: "DMLLatency",
	},
	{
		cwName:         "DMLThroughput",
		prometheusName: "aws_rds_dml_throughput",
		prometheusHelp: "DMLThroughput",
	},
	{
		cwName:         "DatabaseConnections",
		prometheusName: "aws_rds_database_connections",
		prometheusHelp: "The number of database connections in use. Units: Count",
		statistics:     []string{"Sum", "Average"},
	},
	{
		cwName:         "Deadlocks",
		prometheusName: "aws_rds_deadlocks",
		prometheusHelp: "Deadlocks",
	},
	{
		cwName:         "DeleteLatency",
		prometheusName: "aws_rds_delete_latency",
		prometheusHelp: "DeleteLatency",
	},
	{
		cwName:         "DeleteThroughput",
		prometheusName: "aws_rds_delete_throughput",
		prometheusHelp: "DeleteThroughput",
	},
	{
		cwName:         "DiskQueueDepth",
		prometheusName: "aws_rds_disk_queue_depth",
		prometheusHelp: "The number of outstanding IOs (read/write requests) waiting to access the disk. Units: Count",
	},
	{
		cwName:         "EngineUptime",
		prometheusName: "node_boot_time_seconds",
		prometheusHelp: "EngineUptime",
	},
	{
		cwName:         "FreeLocalStorage",
		prometheusName: "aws_rds_free_local_storage",
		prometheusHelp: "FreeLocalStorage",
	},
	{
		cwName:         "FreeStorageSpace",
		prometheusName: "node_filesystem_free_bytes",
		prometheusHelp: "The amount of available storage space. Units: Bytes",
	},
	{
		cwName:         "FreeableMemory",
		prometheusName: "node_memory_Cached_bytes",
		prometheusHelp: "The amount of available random access memory. Units: Bytes",
	},
	{
		cwName:         "InsertLatency",
		prometheusName: "aws_rds_insert_latency",
		prometheusHelp: "InsertLatency",
	},
	{
		cwName:         "InsertThroughput",
		prometheusName: "aws_rds_insert_throughput",
		prometheusHelp: "InsertThroughput",
	},
	{
		cwName:         "LoginFailures",
		prometheusName: "aws_rds_login_failures",
		prometheusHelp: "LoginFailures",
	},
	{
		cwName:         "NetworkReceiveThroughput",
		prometheusName: "aws_rds_network_receive_throughput",
		prometheusHelp: "The incoming (Receive) network traffic on the DB instance, including both customer database traffic and Amazon RDS traffic used for monitoring and replication. Units: Bytes/second",
	},
	{
		cwName:         "NetworkThroughput",
		prometheusName: "aws_rds_network_throughput",
		prometheusHelp: "NetworkThroughput",
	},
	{
		cwName:         "NetworkTransmitThroughput",
		prometheusName: "aws_rds_network_transmit_throughput",
		prometheusHelp: "The outgoing (Transmit) network traffic on the DB instance, including both customer database traffic and Amazon RDS traffic used for monitoring and replication. Units: Bytes/second",
	},
	{
		cwName:         "Queries",
		prometheusName: "aws_rds_queries",
		prometheusHelp: "Queries",
	},
	{
		cwName:         "ReadIOPS",
		prometheusName: "aws_rds_read_iops",
		prometheusHelp: "The average number of disk I/O operations per second. Units: Count/Second",
	},
	{
		cwName:         "ReadLatency",
		prometheusName: "aws_rds_read_latency",
		prometheusHelp: "The average amount of time taken per disk I/O operation. Units: Seconds",
	},
	{
		cwName:         "ReadThroughput",
		prometheusName: "aws_rds_read_throughput",
		prometheusHelp: "The average number of bytes read from disk per second. Units: Bytes/Second",
	},
	{
		cwName:         "ResultSetCacheHitRatio",
		prometheusName: "aws_rds_result_set_cache_hit_ratio",
		prometheusHelp: "ResultSetCacheHitRatio",
	},
	{
		cwName:         "SelectLatency",
		prometheusName: "aws_rds_select_latency",
		prometheusHelp: "SelectLatency",
	},
	{
		cwName:         "SelectThroughput",
		prometheusName: "aws_rds_select_throughput",
		prometheusHelp: "SelectThroughput",
	},
	{
		cwName:         "SwapUsage",
		prometheusName: "aws_rds_swap_usage",
		prometheusHelp: "The amount of swap space used on the DB instance. Units: Bytes",
	},
	{
		cwName:         "UpdateLatency",
		prometheusName: "aws_rds_update_latency",
		prometheusHelp: "UpdateLatency",
	},
	{
		cwName:         "UpdateThroughput",
		prometheusName: "aws_rds_update_throughput",
		prometheusHelp: "UpdateThroughput",
	},
	{
		cwName:         "VolumeBytesUsed",
		prometheusName: "aws_rds_volume_bytes_used",
		prometheusHelp: "VolumeBytesUsed",
	},
	{
		cwName:         "VolumeReadIOPs",
		prometheusName: "aws_rds_volume_read_io_ps",
		prometheusHelp: "VolumeReadIOPs",
	},
	{
		cwName:         "VolumeWriteIOPs",
		prometheusName: "aws_rds_volume_write_io_ps",
		prometheusHelp: "VolumeWriteIOPs",
	},
	{
		cwName:         "WriteIOPS",
		prometheusName: "aws_rds_write_iops",
		prometheusHelp: "The average number of disk I/O operations per second. Units: Count/Second",
	},
	{
		cwName:         "WriteLatency",
		prometheusName: "aws_rds_write_latency",
		prometheusHelp: "The average amount of time taken per disk I/O operation. Units: Seconds",
	},
	{
		cwName:         "WriteThroughput",
		prometheusName: "aws_rds_write_throughput",
		prometheusHelp: "The average number of bytes written to disk per second. Units: Bytes/Second",
	},
	{
		cwName:         "ReplicaLag",
		prometheusName: "aws_rds_replica_lag",
		prometheusHelp: "The amount of time a read replica DB instance lags behind the source DB instance. Unit: Seconds",
	},
}
