package controllers

import (
	"github.com/astaxie/beego"
	"myproject/myblog/models"
	"strconv"
)

type EditController struct {
	beego.Controller
}

func (c *EditController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Params[":id"])
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["Post"] = models.GetBlog(id)
	c.Layout = "layout.tpl"
	c.TplNames = "edit.tpl"
}

func (c *EditController) Post() {
	inputs := c.Input()
	var blog models.Blog
	blog.Id, _ = strconv.Atoi(inputs.Get("Id"))
	blog.Name = c.Input().Get("Name")
	blog.Sex = c.Input().Get("Sex")
	blog.Address = c.Input().Get("Address")
	blog.Title = c.Input().Get("Title")
	blog.Content = c.Input().Get("Content")
	models.SaveBlog(blog)

	c.Ctx.Redirect(302, "/")
}
