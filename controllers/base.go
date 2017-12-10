package controllers

import (
	"github.com/cheneylew/shadowsocks-cms/conf"
	"github.com/cheneylew/goutil/utils"
	beego2 "github.com/cheneylew/goutil/utils/beego"
	"github.com/cheneylew/shadowsocks-cms/models"
)

var FILTER_PATHS  = make([]string, 0)

func init() {
	FILTER_PATHS = append(FILTER_PATHS,"/user/login")
	FILTER_PATHS = append(FILTER_PATHS,"/user/regist")
}

type BaseController struct {
	beego2.BBaseController
}

func (c *BaseController) Prepare() {
	c.Controller.Prepare()
	urlPath := c.Ctx.Request.URL.Path

	c.Data["Website"] = conf.MY_APP_CONFIG.SiteName
	c.Data["Email"] = conf.MY_APP_CONFIG.Email

	c.Layout = "layout.html"

	if !utils.Contain(FILTER_PATHS, urlPath) {
		if c.IsLogin() {
			c.Data["User"] = c.GetUser()
		} else {
			c.RedirectWithURL("/user/login")
		}
	}

}

func (c *BaseController) GetLoginedUser() *models.User {
	v, ok := c.GetUser().(*models.User)
	if ok {
		return v
	}

	return nil
}


