package main

import (
	"github.com/astaxie/beego"
	"github.com/cheneylew/goutil/utils"
	"fmt"
)

func InitTemplateFuncs()  {
	beego.AddFuncMap("flow",utils.FormatFlow)
	beego.AddFuncMap("fmt",Format)
}

func Format(a float64, format string) string {
	return fmt.Sprintf(format, a)
}
