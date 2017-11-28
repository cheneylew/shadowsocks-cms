package conf

type APP_CONFIG struct {
	SiteName string
	Email string
}

var MY_APP_CONFIG APP_CONFIG = APP_CONFIG{
	SiteName:"SS",
	Email:"cheneylew@163.com",
}