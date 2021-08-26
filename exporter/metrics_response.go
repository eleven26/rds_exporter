package exporter

import (
	"encoding/json"
	"fmt"
	"log"
)

type MetricsResponse struct {
	apiResponse ApiResponse
}

func (mp *MetricsResponse) Html() string {
	mappings := mp.apiResponse.Mappings()

	comDML := fmt.Sprintf(`# HELP The average number of Select statements executed per second
# TYPE mysql_com_select gauge
mysql_com_select %s
# HELP The average number of Update statements executed per second
# TYPE mysql_com_update gauge
mysql_com_update %s
# HELP The average number of Delete statements executed per second
# TYPE mysql_com_delete gauge
mysql_com_delete %s
# HELP The average number of Insert statements executed per second
# TYPE mysql_com_insert gauge
mysql_com_insert %s
# HELP The average number of Insert_Select statements executed per second
# TYPE mysql_com_insert_select gauge
mysql_com_insert_select %s
# HELP The average number of Replace statements executed per second
# TYPE mysql_com_replace gauge
mysql_com_replace %s
# HELP The average number of Replace_Select statements executed per second
# TYPE mysql_com_replace_select gauge
mysql_com_replace_select %s
`, mappings["com_select"], mappings["com_update"], mappings["com_delete"], mappings["com_insert"], mappings["com_insert_select"], mappings["com_replace"], mappings["com_replace_select"])

	qpsTps := fmt.Sprintf(`# HELP The average number of transactions per second.
# TYPE mysql_tps gauge
mysql_tps %s
# HELP The average number of SQL statements executed per second
# TYPE mysql_qps gauge
mysql_qps %s
`, mappings["TPS"], mappings["QPS"])

	innoDBBufferRatio := fmt.Sprintf(`# HELP The read hit rate of the InnoDB buffer pool
# TYPE mysql_ibuf_read_hit gauge
mysql_ibuf_read_hit %s
# HELP The usage rate of the InnoDB buffer pool
# TYPE mysql_ibuf_use_ratio gauge
mysql_ibuf_use_ratio %s
# HELP The percentage of dirty blocks in the InnoDB buffer pool.
# TYPE mysql_ibuf_dirty_ratio gauge
mysql_ibuf_dirty_ratio %s
`, mappings["ibuf_read_hit"], mappings["ibuf_use_ratio"], mappings["ibuf_dirty_ratio"])

	tmpDiskTableCreates := fmt.Sprintf(`# HELP The number of temporary tables that are automatically created on the hard disk when the statement is executed.
# TYPE mysql_tb_tmp_disk gauge
mysql_tb_tmp_disk %s
`, mappings["tb_tmp_disk"])

	networkTraffic := fmt.Sprintf(`# HELP The average input traffic per second of the instance. The unit is KB.
# TYPE mysql_recv_k gauge
mysql_recv_k %s
# HELP The average output traffic per second of the instance. The unit is KB.
# TYPE mysql_sent_k gauge
mysql_sent_k %s
`, mappings["recv_k"], mappings["sent_k"])

	innoDBDataReadWriten := fmt.Sprintf(`# HELP The average amount of data that InnoDB reads per second. The unit is KB.
# TYPE mysql_inno_data_read gauge
mysql_inno_data_read %s
# HELP The average amount of data that InnoDB writes per second. The unit is KB.
# TYPE mysql_inno_data_written gauge
mysql_inno_data_written %s
`, mappings["inno_data_read"], mappings["inno_data_written"])

	myISAMKeyReadWrites := fmt.Sprintf(`# HELP The average number of MyISAM reads per second from the buffer pool.
# TYPE mysql_myisam_keyr_r gauge
mysql_myisam_keyr_r %s
# HELP The average number of MyISAM writes per second from the buffer pool.
# TYPE mysql_myisam_keyr_w gauge
mysql_myisam_keyr_w %s
# HELP The average number of MyISAM reads from the hard disk per second.
# TYPE mysql_myisam_keyr gauge
mysql_myisam_keyr %s
# HELP The average MyISAM writes from the hard disk per second The number of entries.
# TYPE mysql_myisam_keyw gauge
mysql_myisam_keyw %s
`, mappings["myisam_keyr_r"], mappings["myisam_keyr_w"], mappings["myisam_keyr"], mappings["myisam_keyw"])

	rowDML := fmt.Sprintf(`# HELP The average number of rows read from the InnoDB table per second.
# TYPE mysql_inno_row_readed gauge
mysql_inno_row_readed %s
# HELP The average number of rows updated from the InnoDB table per second.
# TYPE mysql_inno_row_update gauge
mysql_inno_row_update %s
# HELP The average number of rows deleted from the InnoDB table per second.
# TYPE mysql_inno_row_delete gauge
mysql_inno_row_delete %s
# HELP The average number of rows inserted from the InnoDB table per second.
# TYPE mysql_inno_row_insert gauge
mysql_inno_row_insert %s
# HELP The average number of logs per second The number of physical writes of the file.
# TYPE mysql_inno_log_writes gauge
mysql_inno_log_writes %s
`, mappings["inno_row_readed"], mappings["inno_row_update"], mappings["inno_row_delete"], mappings["inno_row_insert"], mappings["Inno_log_writes"])

	innoDBLogRequests := fmt.Sprintf(`# HELP The average number of reads per second to the InnoDB buffer pool.
# TYPE mysql_ibuf_request_r gauge
mysql_ibuf_request_r %s
# HELP The average number of writes per second to the InnoDB buffer pool.
# TYPE mysql_ibuf_request_w gauge
mysql_ibuf_request_w %s
`, mappings["ibuf_request_r"], mappings["ibuf_request_w"])

	iops := fmt.Sprintf(`# HELP The IOPS (number of IO requests per second) of the instance.
# TYPE mysql_io gauge
mysql_io %s
`, mappings["io"])

	slavestat := fmt.Sprintf(`# HELP Read-only instances delayed time.
# TYPE mysql_iothread gauge
mysql_iothread %s
# TYPE mysql_sqlthread gauge
mysql_sqlthread %s
# TYPE mysql_slave_lag gauge
mysql_slave_lag %s
`, mappings["iothread"], mappings["sqlthread"], mappings["slave-lag"])

	myISAMKeyBufferRatio := fmt.Sprintf(`# HELP MyISAM's average Key Buffer usage per second.
# TYPE mysql_key_usage_ratio gauge
mysql_key_usage_ratio %s
# HELP MyISAM's average Key Buffer read hit rate per second.
# TYPE mysql_key_read_hit_ratio gauge
mysql_key_read_hit_ratio %s
# HELP MyISAM's average Key Buffer write hit rate per second.
# TYPE mysql_key_write_hit_ratio gauge
mysql_key_write_hit_ratio %s
`, mappings["Key_usage_ratio"], mappings["Key_read_hit_ratio"], mappings["Key_write_hit_ratio"])

	innoDBLogWrites := fmt.Sprintf(`# HELP The average number of log write requests per second.
# TYPE mysql_innodb_log_write_requests gauge
mysql_innodb_log_write_requests %s
# HELP The average number of physical writes to the log file per second.
# TYPE mysql_innodb_log_writes gauge
mysql_innodb_log_writes %s
# HELP The average number of fsync() writes to the log file per second.
# TYPE mysql_innodb_os_log_fsyncs gauge
mysql_innodb_os_log_fsyncs %s
`, mappings["Innodb_log_write_requests"], mappings["Innodb_log_writes"], mappings["Innodb_os_log_fsyncs"])

	memCpuUsage := fmt.Sprintf(`# HELP Instance CPU usage (accounting for the total number of operating systems).
# TYPE mysql_cpuusage gauge
mysql_cpuusage %s
# HELP MySQL instance memory usage (accounting for the total number of operating systems).
# TYPE mysql_memusage gauge
mysql_memusage %s
`, mappings["cpuusage"], mappings["memusage"])

	sessions := fmt.Sprintf(`# HELP Current number of active connections.
# TYPE mysql_active_session gauge
mysql_active_session %s
# HELP Current total number of connections.
# TYPE mysql_total_session gauge
mysql_total_session %s
`, mappings["active_session"], mappings["total_session"])

	threadStatus := fmt.Sprintf(`# HELP The number of active threads.
# TYPE mysql_threads_running gauge
mysql_threads_running %s
# HELP The number of thread connections.
# TYPE mysql_threads_connected gauge
mysql_threads_connected %s
`, mappings["threads_running"], mappings["threads_connected"])

	detailedSpaceUsage := fmt.Sprintf(`# HELP The total instance space usage.
# TYPE mysql_ins_size gauge
mysql_ins_size %s
# HELP The total data space usage.
# TYPE mysql_data_size gauge
mysql_data_size %s
# HELP The total log space usage.
# TYPE mysql_log_size gauge
mysql_log_size %s
# HELP The total temporary space usage.
# TYPE mysql_tmp_size gauge
mysql_tmp_size %s
# HELP The total system space usage.
# TYPE mysql_other_size gauge
mysql_other_size %s
`, mappings["ins_size"], mappings["data_size"], mappings["log_size"], mappings["tmp_size"], mappings["other_size"])

	return comDML + qpsTps + innoDBBufferRatio + tmpDiskTableCreates + networkTraffic + innoDBDataReadWriten + myISAMKeyReadWrites + rowDML + innoDBLogRequests + iops + slavestat + myISAMKeyBufferRatio + innoDBLogWrites + memCpuUsage + sessions + threadStatus + detailedSpaceUsage
}

func (mp *MetricsResponse) Raw() []Metrics {
	return mp.apiResponse.Metrics()
}

func (mp *MetricsResponse) Json() string {
	metrics := mp.apiResponse.Mappings()

	res, err := json.Marshal(metrics)
	if err != nil {
		log.Fatal(err)
	}

	return string(res)
}
