package service

import (
	"hello/models"
)

type UserService struct {
	BaseService
}

func (s *UserService)GetUser(user *models.Users, userId int) {
	O.QueryTable(new(models.Users)).Filter("id",userId).One(user)
}

func (s *UserService)GetUserOrders(list *[]models.GameBetting, userId, gameId, limit int){
	O.QueryTable(new(models.GameBetting)).Filter("user_id",userId).Filter("game_id",gameId).Limit(limit).OrderBy("-betting_time").All(list)
}
