package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	//退出操作
	fmt.Println("进入login")
	if c.Input().Get("exit") == "true" {
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("userpwd", "", -1, "/")
		c.Redirect("/login", 302) //////////////////////////////////退出跳转至登录界面//////////////////////////////
		return
	}

	c.TplNames = "login.tpl"
}

func (c *LoginController) Post() {
	//表单信息
	username := c.Input().Get("username")
	userpwd := c.Input().Get("userpwd")
	autoLogin := c.Input().Get("autoLogin") == "on"

	//验证用户和密码

	if username == beego.AppConfig.String("username") && userpwd == beego.AppConfig.String("userpwd") {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("username", username, maxAge, "/")
		c.Ctx.SetCookie("userpwd", userpwd, maxAge, "/")
		c.Redirect("/", 302) ////////////////////////////////////密码和用户成功后执行///////////////////////

	}
	c.Redirect("/login", 302) ////////////////////////////////////密码和用户失败后执行///////////////////////
	return
}

func checkAccount(ctx *context.Context) bool {
	fmt.Println("检查账号")
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value

	ck, err = ctx.Request.Cookie("userpwd")
	if err != nil {
		return false
	}
	userpwd := ck.Value

	return username == beego.AppConfig.String("username") && userpwd == beego.AppConfig.String("userpwd")

}
