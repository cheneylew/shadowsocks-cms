package controllers

import (
	"github.com/cheneylew/goutil/utils"
	"github.com/cheneylew/shadowsocks-cms/database"
	"github.com/cheneylew/shadowsocks-cms/models"
	"time"
)

type UserController struct {
	BaseController
}

func (c *UserController) Prepare() {
	c.BaseController.Prepare()
}

func (c *UserController) Finish() {
	c.BaseController.Finish()
}

func (c *UserController) Get() {
	c.TplName = "user_login.html"
}

func (c *UserController) Login() {
	c.TplName = "user_login.html"
	email := c.GetString("email")
	password := c.GetString("password")
	utils.JJKPrintln(email, password)

	if len(email) > 0 && len(password) > 0 {
		users := database.DBQueryUserWithEmailOrMobile(email)

		var loginedUser models.User
		isLogin := false
		for _, user := range users {
			if user.Password == password {
				loginedUser = user
				isLogin = true
			}
		}

		if isLogin {
			c.SetLoginedUser(loginedUser)
			c.RedirectWithURL("/user/home")
		}
	}
}

func (c *UserController) Regist() {
	c.TplName = "user_regist.html"
}

func (c *UserController) Logout() {
	c.SetUserLogout()
	c.RedirectWithURL("/user/login")
}

func (c *UserController) Home() {
	c.TplName = "user_home.html"

	utils.JJKPrintln()
	ports := database.DBQueryPortsWithUserId(c.GetLoginedUser().User_id)
	for i := 0; i < len(ports); i++ {
		ports[i].Flow_surplus = ports[i].Flow_in_max - ports[i].Flow_in
		days := float64(ports[i].End_time.UTC().Unix()-time.Now().UTC().Unix())/(float64(24*60*60))
		if days > 0 {
			ports[i].Days_surplus = days
		} else {
			ports[i].Days_surplus = 0
		}
	}
	c.Data["Ports"] = ports
}

func (c *UserController) Setting() {
	c.TplName = "user_setting.html"

	if c.IsGet() {
		return
	}

	refer := c.GetString("refer")
	comment := c.GetString("comment")
	password := c.GetString("password")

	user := c.GetLoginedUser()
	user.Refer = refer
	user.Comment = comment
	user.Password = password
	n, err := database.O.Update(user,"Refer","Comment","Password")
	if err != nil {
		utils.JJKPrintln(err)
	}
	if n > 0 {
		c.SetLoginedUser(*user)
		c.Data["User"] = user
	}
}



