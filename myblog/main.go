package main

import (
	"github.com/astaxie/beego"
	_ "myproject/myblog/routers"
)

func main() {
	beego.Run()
}
