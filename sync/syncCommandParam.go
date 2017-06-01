package sync




type SyncCmdParam struct {

	Command string	// python2.6 python
	DataxPath string	// datax 路径
	Mode string	// 同步模式 mysql2hdfs  mysql2mysql
	EnableNotify string // 是否通知业务系统
	NotifyUrl string // 通知业务系统url
}



