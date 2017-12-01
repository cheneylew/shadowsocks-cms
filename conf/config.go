package conf

type APP_CONFIG struct {
	SiteName string
	Email string
}

var MY_APP_CONFIG APP_CONFIG = APP_CONFIG{
	SiteName:"流量统计",
	Email:"cheneylew@163.com",
}