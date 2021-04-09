package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "hello/models"
	_ "hello/routers"
	_ "hello/service"
)

func main() {
	web.Run()
}

