package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myproject/myblog/models"
	"strconv"
)

type ViewController struct {
	beego.Controller
}

func (c *ViewController) Get() {

	id, _ := strconv.Atoi(c.Ctx.Input.Params[":id"])
	fmt.Println("id", id)

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["Blogs"] = models.GetBlog(id)
	fmt.Println("models.GetBlog(id) =:", c.Data["Blogs"])
	c.Layout = "layout.tpl"
	c.TplNames = "view.tpl"
}
