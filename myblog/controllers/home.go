package controllers

import (
	"github.com/astaxie/beego"
	"myproject/myblog/models"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["Blogs"] = models.GetAll()
	c.Layout = "layout.tpl"
	c.TplNames = "home.tpl"

}
