package controllers

import (
	"github.com/astaxie/beego"
	"myproject/myblog/models"
	"strconv"
)

type AddController struct {
	beego.Controller
}

func (c *AddController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Layout = "layout.tpl"
	c.TplNames = "add.tpl"
}

func (c *AddController) Post() {
	inputs := c.Input()
	var blog models.Blog
	blog.Id, _ = strconv.Atoi(inputs.Get("Id"))
	blog.Name = inputs.Get("Name")
	blog.Sex = inputs.Get("Sex")
	blog.Address = inputs.Get("Address")
	blog.Title = c.Input().Get("Title")
	blog.Content = c.Input().Get("Content")
	models.AddBlog(blog, blog.Id)
	c.Ctx.Redirect(302, "/")
}
