package main

import (
	_ "my-first-beego-project/routers"

	"my-first-beego-project/controllers"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
