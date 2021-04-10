package controller

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"hello/jobs/rabbitMq/rabbitmq"
	"hello/jobs/resultEntry/model"
	"time"
)

var O orm.Ormer

func init(){
	O = orm.NewOrm()
}

func ResultEntry(){
	//var gamePlaySlice = make([]model.GamePlay, 10)
	var gamePlaySlice []orm.ParamsList
	O.QueryTable(new(model.GamePlay)).Filter("end_time__gte",time.Now().Unix()).Filter("status",0).OrderBy("-end_time").Limit(10).ValuesList(&gamePlaySlice,"id","game_id","number")
	rabbitmq := rabbitmq.NewRabbitMQSimple("result_game_AUTO")
	for _,gamePlay := range gamePlaySlice{
		per,_ := json.Marshal(gamePlay)
		rabbitmq.PublishSimple(string(per))
	}
}