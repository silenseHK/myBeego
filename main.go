package main

import (
	"github.com/astaxie/beego"
	_ "hello/models"
	_ "hello/routers"
	_ "hello/service"
)

func main() {
	beego.Run()
}

