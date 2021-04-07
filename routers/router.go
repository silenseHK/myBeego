package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"hello/controllers"
	"hello/controllers/api"
)


func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/game", &api.GameController{})
    //beego.Router("/game2", &api.GameController{},"get:GameStart")

    gamePath := beego.AppConfig.String("gamePath")
    fmt.Println(gamePath)
    ns :=
    	beego.NewNamespace("/" + gamePath,
    			beego.NSRouter("/", &api.GameController{}),
    			beego.NSRouter("/game_start", &api.GameController{},"get:GameStart"),
    		)
    beego.AddNamespace(ns)
}
