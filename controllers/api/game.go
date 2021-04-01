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

var (
	gameService *service.GameService
	userService *service.UserService
)

func  (c *GameController) Prepare(){
	gameService = new(service.GameService)
	userService = new(service.UserService)
}


/*
	获取本期数据
 */
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
		gameService.GetGameStart(&gamePlay, gameId)
	}()
	//用户数据
	wg.Add(1)
	go func(){
		defer wg.Done()
		userService.GetUser(&user, userId)
	}()
	//上期数据
	wg.Add(1)
	go func(){
		defer wg.Done()
		gameService.GetPreGame(&preGamePlay, gameId)
	}()
	//往十期的结果
	wg.Add(1)
	go func(){
		defer wg.Done()
		gameService.GetPreGameList(&preList, gameId,1,10)
	}()
	//用户的投注订单
	wg.Add(1)
	go func(){
		defer wg.Done()
		userService.GetUserOrders(&orderList,userId,gameId,5)
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

func (c *GameController)GameStart(){  //测试协程和不协程的请求时间
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
	gs := new(service.GameService)
	gs.GetGameStart(&gamePlay, gameId)
	//用户数据
	us := new(service.UserService)
	us.GetUser(&user, userId)
	//上期数据
	gs2 := new(service.GameService)
	gs2.GetPreGame(&preGamePlay, gameId)
	//往十期的结果
	new(service.GameService).GetPreGameList(&preList, gameId,1,10)
	//用户的投注订单
	new(service.UserService).GetUserOrders(&orderList,userId,gameId,5)

	c.ReturnJsonWithData(200,"", map[string]interface{}{
		"game_play": gamePlay,
		"user": user,
		"pre_game_play": preGamePlay,
		"pre_list": preList,
		"order_list": orderList,
	})
}

func (c *GameController)Betting(){
	//userId,err := c.GetInt("user_id")
	//if err != nil{
	//	c.ReturnJsonWithData(300,"参数缺失","")
	//	return
	//}
	//gameId,err := c.GetInt("game_id")
	//if err != nil{
	//	c.ReturnJsonWithData(300,"参数缺失","")
	//	return
	//}
	//gameConfigId,err := c.GetInt("game_c_id")
	//if err != nil{
	//	c.ReturnJsonWithData(300,"参数缺失","")
	//	return
	//}
	//bettingMoney,err := c.GetFloat("betting_money")
	//if err != nil{
	//	c.ReturnJsonWithData(300,"参数缺失","")
	//	return
	//}
}

func (c *GameController) PushGame(){

}
