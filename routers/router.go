package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"hello/controllers"
	"hello/controllers/api"
	"hello/middlewares"
)


func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/game", &api.GameController{})
    //beego.Router("/game2", &api.GameController{},"get:GameStart")

    gamePath := beego.AppConfig.String("gamePath")
    fmt.Println(gamePath)
    ns :=
    	beego.NewNamespace("/" + gamePath,
    			beego.NSBefore(middlewares.CheckUserToken),
    			beego.NSRouter("/", &api.GameController{}),
    			beego.NSRouter("/start", &api.GameController{},"post:GameStart"),
    		)
    beego.AddNamespace(ns)
}
