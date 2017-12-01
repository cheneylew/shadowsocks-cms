package database

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/cheneylew/goutil/utils"
	"fmt"
	"time"
	"github.com/cheneylew/shadowsocks-cms/models"
)

var o orm.Ormer

func init() {
	url := dbUrl("cheneylew","12344321","47.91.151.207","3308","shadowsocks-servers")
	utils.JJKPrintln(url)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", url)
	o = orm.NewOrm()

	if err != nil {
		utils.JJKPrintln("========database can't connect! error:" + err.Error()+"========")
	} else {
		utils.JJKPrintln("========database connected success！========")
	}

}

func dbUrl(user, password, host, port, dbName string) string {
	return fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8`, user, password, host, port, dbName)
}

func DBQueryServers(ip string) []*models.Server {
	var objects []*models.Server

	qs := o.QueryTable("server")
	_, err := qs.Filter("ip", ip).All(&objects)

	if err != nil {
		return objects
	}

	return objects
}

func DBQueryPortsWithSid(sid int64) []*models.Port {
	var objects []*models.Port

	qs := o.QueryTable("port")
	_, err := qs.Filter("server__server_id", sid).RelatedSel().All(&objects)

	if err != nil {
		return objects
	}

	return objects
}

func DBQueryPortsWithUserId(userId int64) []*models.Port {
	var objects []*models.Port

	qs := o.QueryTable("port")
	_, err := qs.Filter("user__user_id", userId).RelatedSel().All(&objects)

	if err != nil {
		return objects
	}

	return objects
}

func DBQueryPortsWithIP(ip string) []*models.Port {
	var objects []*models.Port

	qs := o.QueryTable("port")
	_, err := qs.Filter("server__ip", ip).RelatedSel().All(&objects)

	if err != nil {
		return objects
	}

	return objects
}

func DBQueryUsersWithUid(uid int64) []*models.User {
	var objects []*models.User
	qs := o.QueryTable("user")
	_, err := qs.Filter("user_id", uid).RelatedSel().All(&objects)

	if err != nil {
		return objects
	}

	return objects
}

func DBQueryUserWithEmailOrMobile(emailOrMobile string) []models.User {
	var objects []models.User
	_, err := o.Raw(fmt.Sprintf("select * from user where email =? or mobile=?"),emailOrMobile, emailOrMobile).QueryRows(&objects)
	if err != nil {
		utils.JJKPrintln(err)
		return objects
	}

	return objects
}

func DBQueryMyListenPorts() []*models.Port {
	curIp, _ := utils.ExtranetIP()
	ports := DBQueryPortsWithIP(curIp)

	var filterPorts []*models.Port
	for _, port := range ports {
		if port.Ptype == 0 {
			//包年包月
			//判断截止时间
			if time.Now().Before(port.End_time) {
				filterPorts = append(filterPorts, port)
			}
		} else if port.Ptype == 1 {
			//包流量
			//流量是否超限
			if port.Flow_total < port.Flow_in_max {
				filterPorts = append(filterPorts, port)
			}
		}
	}


	return filterPorts
}

func dbStart()  {
	o = orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库


}