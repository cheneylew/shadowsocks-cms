package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cheneylew/shadowsocks-cms/conf"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	c.Controller.Prepare()
	c.Data["Website"] = conf.MY_APP_CONFIG.SiteName
	c.Data["Email"] = conf.MY_APP_CONFIG.Email

	c.Layout = "layout.html"
}

func (c *MainController) Finish() {
	c.Controller.Finish()
}

func (c *MainController) Get() {
	c.TplName = "user_login.html"
}

// @router /user/login/ [get]
func (c *MainController) UserLogin() {
	c.TplName = "user_login.html"
}

// @router /user/regist/ [get]
func (c *MainController) UserRegist() {
	c.TplName = "user_regist.html"
}


