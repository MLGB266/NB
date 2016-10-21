package controllers

import (
	"github.com/astaxie/beego"
	"myproject/myblog/models"
	"strconv"
)

type DeleteController struct {
	beego.Controller
}

func (c *DeleteController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Params[":id"])
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	models.DelBlog(id)
	c.Ctx.Redirect(302, "/")
}
