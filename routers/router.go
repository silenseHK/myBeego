package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"hello/controllers"
	"hello/controllers/api"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/game", &api.GameController{})
}
