package controller

import (
	"github.com/beego/beego/v2/client/orm"
	"hello/jobs/generateNumber/model"
	"time"
)

var O orm.Ormer

func init(){
	O = orm.NewOrm()
}

func ResultEntry(){
	var gamePlaySlice = make([]model.GamePlay, 10)
	O.QueryTable(new(model.GamePlay)).Filter("end_time__lte",time.Now().Unix()).Filter("status",0).OrderBy("-end_time").Limit(10).All(&gamePlaySlice)
}