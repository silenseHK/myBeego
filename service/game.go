package service

import (
	"hello/models"
	"time"
)

type GameService struct{
	BaseService
}

func (s GameService)GetGameStart(gamePlay *models.GamePlay, gameId int){
	rightNow := time.Now().Unix()
	O.QueryTable(new(models.GamePlay)).Filter("game_id", gameId).Filter("status",0).Filter("end_time__gte", rightNow).OrderBy("end_time").One(gamePlay)  //如果查询不到就会返回error
}

func (s GameService)GetPreGame(gamePlay *models.GamePlay, gameId int){
	rightNow := time.Now().Unix()
	O.QueryTable(new(models.GamePlay)).Filter("game_id", gameId).Filter("end_time__lt", rightNow).OrderBy("-end_time").One(gamePlay,"Id","PrizeNumber")
}

func (s GameService)GetPreGameList(list *[]models.GamePlay, gameId, offset, limit int){
	rightNow := time.Now().Unix()
	O.QueryTable(new(models.GamePlay)).Filter("game_id", gameId).Filter("end_time__lt", rightNow).OrderBy("-end_time").Offset(offset).Limit(limit).All(list,"Id","PrizeNumber","Number")
}
