package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cheneylew/shadowsocks-cms/conf"
	"github.com/cheneylew/shadowsocks-cms/models"
	"github.com/cheneylew/goutil/utils"
)

const SESSTION_KEY_USER  = "LOGINED_USER"
var FILTER_PATHS  = make([]string, 0)

func init() {
	FILTER_PATHS = append(FILTER_PATHS,"/user/login")
}

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.Controller.Prepare()
	urlPath := c.Ctx.Request.URL.Path

	c.Data["Website"] = conf.MY_APP_CONFIG.SiteName
	c.Data["Email"] = conf.MY_APP_CONFIG.Email

	c.Layout = "layout.html"

	if !utils.Contain(FILTER_PATHS, urlPath) {
		if c.IsLogin() {
			user := c.GetLoginedUser()
			c.Data["User"] = user
			utils.JJKPrintln(user)
		} else {
			c.RedirectWithURL("/user/login")
		}
	}

}

func (c *BaseController) Finish() {
	c.Controller.Finish()
}

//通用

func (c *BaseController) RedirectWithURL(url string) {
	c.Redirect(url, 302)
}

//用户

func (c *BaseController) IsLogin() bool {
	utils.JJKPrintln("=====get ",c.GetSession(SESSTION_KEY_USER))
	return c.GetSession(SESSTION_KEY_USER) != nil
}

func (c *BaseController) SetLoginedUser(user models.User) {
	utils.JJKPrintln("=====set ",user)
	c.SetSession(SESSTION_KEY_USER,user)
}

func (c *BaseController) SetUserLogout() {
	c.SetSession(SESSTION_KEY_USER,nil)
}

func (c *BaseController) GetLoginedUser() *models.User {
	v, b := c.GetSession(SESSTION_KEY_USER).(models.User)
	if b {
		return &v
	}

	return nil
}
