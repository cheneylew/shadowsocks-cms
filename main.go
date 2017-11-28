package main

import (
	"github.com/astaxie/beego"
	_ "github.com/cheneylew/shadowsocks-cms/routers"
)

func main() {
	beego.Run()
}
