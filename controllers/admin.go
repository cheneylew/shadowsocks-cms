package controllers

import (
	"github.com/cheneylew/shadowsocks-cms/database"
	"github.com/cheneylew/goutil/utils"
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
	sid := c.PathValueInt()

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

	c.TplName = "admin_port_add.html"

}
