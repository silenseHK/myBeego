package main

import (
	"github.com/astaxie/beego"
	_ "hello/models"
	_ "hello/routers"
	_ "hello/service"
	_ "hello/tools"
)

func main() {
	beego.Run()
}

