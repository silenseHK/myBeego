package api

import (
	"hello/models"
	"hello/service"
	"sync"
)

type GameController struct {
	BaseApiController
}

var wg sync.WaitGroup

func (c *GameController) Get(){
	gameId,_ := c.GetInt("game_id")
	userId,_ := c.GetInt("user_id")

	var (
		gamePlay models.GamePlay
		user models.Users
		preGamePlay models.GamePlay
		preList []models.GamePlay
		orderList []models.GameBetting
	)

	//本期数据
	wg.Add(1)
	go func(){
		defer wg.Done()
		gs := new(service.GameService)
		gs.GetGameStart(&gamePlay, gameId)
	}()
	//用户数据
	wg.Add(1)
	go func(){
		defer wg.Done()
		us := new(service.UserService)
		us.GetUser(&user, userId)
	}()
	//上期数据
	wg.Add(1)
	go func(){
		defer wg.Done()
		gs := new(service.GameService)
		gs.GetPreGame(&preGamePlay, gameId)
	}()
	//往十期的结果
	wg.Add(1)
	go func(){
		defer wg.Done()
		new(service.GameService).GetPreGame10(&preList, gameId)
	}()
	//用户的投注订单
	wg.Add(1)
	go func(){
		defer wg.Done()
		new(service.UserService).GetUserOrders(&orderList,userId,gameId,10)
	}()

	wg.Wait()
	c.ReturnJsonWithData(200,"", map[string]interface{}{
		"game_play": gamePlay,
		"user": user,
		"pre_game_play": preGamePlay,
		"pre_list": preList,
		"order_list": orderList,
	})
}
