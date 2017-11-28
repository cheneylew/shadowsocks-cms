package routers

import (
	"github.com/astaxie/beego"
	"github.com/cheneylew/shadowsocks-cms/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.MainController{})
}
