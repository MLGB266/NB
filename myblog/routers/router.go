package routers

import (
	"github.com/astaxie/beego"
	"myproject/myblog/controllers"
)

func init() {
	//完成home页
	beego.Router("/", &controllers.HomeController{})

	//完成login页
	beego.Router("/login", &controllers.LoginController{})

	//完成view页
	beego.Router("/view/:id([0-9]+)", &controllers.ViewController{})

	//完成edit页
	beego.Router("/edit/:id([0-9]+)", &controllers.EditController{})

	//完成delete页
	beego.Router("/delete/:id([0-9]+)", &controllers.DeleteController{})

	//完成add页面
	beego.Router("/add", &controllers.AddController{})

}
