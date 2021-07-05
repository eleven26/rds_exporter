package exporter

import "strings"

// 性能参数表 https://help.aliyun.com/document_detail/26316.html
var mappings = []string{
	"MySQL_NetworkTraffic",
	"MySQL_QPSTPS",
	"MySQL_Sessions",
	"MySQL_InnoDBBufferRatio",
	"MySQL_InnoDBDataReadWriten",
	"MySQL_InnoDBLogRequests",
	"MySQL_InnoDBLogWrites",
	"MySQL_TempDiskTableCreates",
	"MySQL_MyISAMKeyBufferRatio",
	"MySQL_MyISAMKeyReadWrites",
	"MySQL_COMDML",
	"MySQL_RowDML",
	"MySQL_MemCpuUsage",
	"MySQL_IOPS",
	"MySQL_DetailedSpaceUsage",
	"slavestat",
	"MySQL_ThreadStatus",
	"MySQL_ReplicationDelay",
	"MySQL_ReplicationThread",
}

func PerformanceKeysString() string {
	return strings.Join(mappings, ",")
}
