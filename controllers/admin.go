package controllers

import (
	"github.com/cheneylew/shadowsocks-cms/database"
	"github.com/cheneylew/goutil/utils"
	"github.com/cheneylew/shadowsocks-cms/models"
	"github.com/jinzhu/now"
	"fmt"
	"time"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) Prepare() {
	c.BaseController.Prepare()
	if !c.GetLoginedUser().Isadmin {
		c.RedirectWithURL("/user/home")
	}
}

func (c *AdminController) Finish() {
	c.BaseController.Finish()
}

func (c *AdminController) Home() {
	c.TplName = "admin_home.html"
	servers := database.DBQueryServersAll()
	utils.JJKPrintln(servers)
	c.Data["Servers"] = servers
}

func (c *AdminController) Ports() {
	sid := c.PathValueInt()

	c.TplName = "admin_ports.html"
	ports := database.DBQueryPortsWithSid(int64(sid))
	for i := 0; i < len(ports); i++ {
		ports[i].Flow_surplus = ports[i].Flow_in_max - ports[i].Flow_in
	}
	c.Data["Ports"] = ports

	server := database.DBQueryServersWithSid(sid)
	c.Data["Server"] = server
}

func (c *AdminController) PortAdd() {
	c.TplName = "admin_port_add.html"
	sid := c.PathValueInt()
	if c.IsGet() {
		server := database.DBQueryServersWithSid(sid)
		c.Data["Server"] = server

		c.Data["Users"] = database.DBQueryUsersAll()
		c.Data["RandomPwd"] = utils.RandomString(32)

		maxPort := database.DBQueryMaxPortWithIP(server.Ip)
		if maxPort != nil {
			c.Data["MaxPortNum"] = utils.JKStrToInt64(maxPort.Port)+1
		} else {
			c.Data["MaxPortNum"] = 10000
		}


	} else {
		endTime, _ := now.Parse(c.GetString("end_time"))
		user := database.DBQueryUserWithUid(utils.JKStrToInt64(c.GetString("user_id")))
		server := database.DBQueryServersWithSid(sid)
		port := &models.Port{
			User:user,
			Server:server,
			Ptype:utils.JKStrToUInt8(c.GetString("ptype")),
			Port:c.GetString("port"),
			Password:c.GetString("password"),
			Flow_in_max:utils.ToFloat64(c.GetString("flow_in_max")),
			End_time:endTime,
		}
		_, e := database.O.Insert(port)
		if e != nil {
			utils.JJKPrintln(e)
		}

		c.RedirectWithURL(fmt.Sprintf("/admin/ports/%d", sid))
	}
}

func (c *AdminController) PortUpdate() {
	c.TplName = "admin_port_update.html"
	pid := c.PathValueInt()
	if c.IsGet() {
		port := database.DBQueryPortWithPid(pid)
		c.Data["Port"] = port
		c.Data["Users"] = database.DBQueryUsersAll()


	} else {
		endTime, _ := now.Parse(c.GetString("end_time"))
		endTime = endTime.Add(time.Hour*8)
		user := database.DBQueryUserWithUid(utils.JKStrToInt64(c.GetString("user_id")))

		port := database.DBQueryPortWithPid(pid)
		port.End_time = endTime
		port.User = user
		port.Ptype = utils.JKStrToUInt8(c.GetString("ptype"))
		port.Port = c.GetString("port")
		port.Password = c.GetString("password")
		port.Flow_in_max = utils.ToFloat64(c.GetString("flow_in_max"))
		port.Flow_in = utils.ToFloat64(c.GetString("flow_in"))

		utils.JJKPrintln(endTime)

		_, e := database.O.Update(port)

		if e != nil {
			utils.JJKPrintln(e)
		}

		c.RedirectWithURL(fmt.Sprintf("/admin/portupdate/%d", pid))
	}
}
