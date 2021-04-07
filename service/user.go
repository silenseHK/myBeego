package service

import (
	"hello/models"
	"time"
)

type UserService struct {
	BaseService
}

func (s *UserService)GetUser(user *models.Users, userId int) error{
	e := O.QueryTable(new(models.Users)).Filter("id",userId).One(user)
	return e
}

func (s *UserService)GetUserOrders(list *[]models.GameBetting, userId, gameId, limit int){
	O.QueryTable(new(models.GameBetting)).Filter("user_id",userId).Filter("game_id",gameId).Limit(limit).OrderBy("-betting_time").All(list)
}

func (s UserService)AddBalanceLog(userId int,dqBalance float64, wcBalance float64, _type int8)(id int64, e error){
	var userBalanceLog = models.UserBalanceLogs{
		UserId: userId,
		DqBalance: dqBalance,
		WcBalance: wcBalance,
		Type: _type,
		Time: time.Now().Unix(),
	}
	id,e = O.Insert(&userBalanceLog)
	return
}
