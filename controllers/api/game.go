package api

import (
	"fmt"
	"github.com/shopspring/decimal"
	"hello/libs"
	"hello/models"
	"hello/service"
	"sync"
	"time"
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
	c.ReturnJson()
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

func (c *GameController)Betting(){
	fmt.Println(c.O)
	userId,err := c.GetInt("UserId")
	if err != nil{
		c.ReturnJsonWithData(300,c.Lang["params"]["miss"],"")
		return
	}
	gameId,err := c.GetInt("game_id")
	if err != nil{
		c.ReturnJsonWithData(300,c.Lang["params"]["miss"],"")
		return
	}
	gameConfigId,err := c.GetInt("game_c_id")
	if err != nil{
		c.ReturnJsonWithData(300,c.Lang["params"]["miss"],"")
		return
	}
	bettingMoney,err := c.GetFloat("betting_money")
	if err != nil{
		c.ReturnJsonWithData(300,c.Lang["params"]["miss"],"")
		return
	}
	user := models.Users{}
	e := userService.GetUser(&user, userId)
	if e != nil{
		c.ReturnJsonWithData(300,e.Error(),"")
		return
	}
	if user.Balance < bettingMoney{
		c.ReturnJsonWithData(300,c.Lang["balance"]["less"],"")
		return
	}
	gamePlay,e := gameService.GetGame(gameId)
	if e != nil{
		c.ReturnJsonWithData(300,e.Error(),"")
		return
	}
	if gamePlay.EndTime + 10 <= time.Now().Unix(){
		c.ReturnJsonWithData(300,c.Lang["game"]["timeOut"],"")
		return
	}
	gameConfig,e := gameService.GetGameConfig(gameConfigId)
	if e != nil{
		c.ReturnJsonWithData(300,e.Error(),"")
		return
	}
	if gameConfig.GameId != gamePlay.GameId{
		c.ReturnJsonWithData(300,c.Lang["game"]["chooseWrong"],"")
		return
	}

	TO,_ := c.O.Begin()
	wcBalance := decimal.NewFromFloat(user.Balance).Sub(decimal.NewFromFloat(bettingMoney))
	wcBalanceF,_ := wcBalance.Float64()
	//增加用户余额变动记录
	userService.AddBalanceLog(userId, user.Balance, wcBalanceF,1)

	//修改用户余额
	user.Balance = wcBalanceF
	if _,e = c.O.Update(&user); e!=nil{
		c.ReturnJsonWithData(300,c.Lang["game"]["balanceDeductionFail"],"")
		TO.Rollback()
		return
	}

	serviceCharge,_ := decimal.NewFromFloat(bettingMoney).Mul(decimal.NewFromFloat(0.97)).Round(2).Float64()
	money,_ := decimal.NewFromFloat(bettingMoney).Sub(decimal.NewFromFloat(serviceCharge)).Float64()
	gameBetting := models.GameBetting{
		BettingNum: gamePlay.Number,
		UserId: userId,
		GameId: gamePlay.GameId,
		GamePId: gamePlay.Id,
		GameCXId: gameConfigId,
		Money: money,
		Odds: gameConfig.Odds,
		BettingTime: time.Now().Unix(),
		ServiceCharge: serviceCharge,
	}
	_,e = c.O.Insert(&gameBetting)
	if e != nil{
		c.ReturnJsonWithData(300,c.Lang["game"]["bettingFail"],"")
		TO.Rollback()
		return
	}
	TO.Commit()
	c.ReturnJsonWithData(libs.Success, c.Lang["game"]["bettingSuccess"],"")
}

func (c *GameController) PushGame(){

}
