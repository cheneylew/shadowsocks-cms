package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/cheneylew/goutil/utils"
)

type Server struct {
	Server_id int64 `orm:"pk"`
	Ip string
}

type Port struct {
	Port_id int64 `orm:"pk"`
	Port string
	Password string
	Server *Server `orm:"rel(one)"`
	User *User `orm:"rel(one)"`
	Flow_in float64
	Flow_out float64
	Flow_total float64
	Flow_in_max float64
	Ptype uint8
	Start_time time.Time
	End_time time.Time
	Sync_time time.Time

	Flow_surplus float64 `orm:"-"`	//剩余流量
	Days_surplus float64 `orm:"-"`	//剩余天数
}

type User struct {
	User_id int64       `orm:"pk"`
	Name string
	Email string
	Mobile string
	Refer string
	Comment string
	Password string
}

func init() {
	utils.JJKPrintln("========regist model=======")
	orm.RegisterModel(new(User), new(Server), new(Port))
}