package routers

import (
	"github.com/beego/beego/v2/server/web"
	"hello/controllers/api"
	"hello/middlewares"
)

func init() {
    gamePath,_ := web.AppConfig.String("gamePath")
    ns :=
    	web.NewNamespace("/" + gamePath,
    			web.NSBefore(middlewares.CheckUserToken),
    			web.NSRouter("/", &api.GameController{}),
    			web.NSRouter("/start", &api.GameController{},"post:GameStart"),
				web.NSRouter("/betting", &api.GameController{},"post:Betting"),
    		)
	web.AddNamespace(ns)

	web.Router("/login",&api.UserController{},"post:Login")
}
