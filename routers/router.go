package routers

import (
	"github.com/astaxie/beego"
	"github.com/cheneylew/shadowsocks-cms/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.MainController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.AdminController{})
}
