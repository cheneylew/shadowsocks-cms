package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user.html"

	c.Layout = "layout.html"
	c.TplName = "user_login.html"
}

// @router /user/login/ [get]
func (c *MainController) UserLogin() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user.html"

	c.Layout = "layout.html"
	c.TplName = "user_login.html"
}

// @router /user/regist/ [get]
func (c *MainController) UserRegist() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user.html"

	c.Layout = "layout.html"
	c.TplName = "user_regist.html"
}


