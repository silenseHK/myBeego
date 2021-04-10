package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/beego/beego/v2/client/orm"
	"github.com/garyburd/redigo/redis"
	"github.com/shopspring/decimal"
	"github.com/streadway/amqp"
	"hello/jobs/rabbitMq/rabbitmq"
	"hello/jobs/rabbitMq/resultNumberAuto/model"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var O orm.Ormer

func init(){
	O = orm.NewOrm()
}

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("result_game_AUTO")
	channel,queue,err := rabbitmq.ConsumeObj()
	if err != nil{
		log.Printf(err.Error())
		return
	}
	forever := make(chan bool)
	msgs, err := channel.Consume(
		queue, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for d := range msgs {
			//消息逻辑处理，可以自行设计逻辑
			go resultNumber(d)
		}
	}()
	log.Printf("Wait...")
	<-forever
}

func resultNumber(d amqp.Delivery){
	var str = strings.Replace(strings.Trim(strings.Trim(string(d.Body),"["),"]"),"\"","",2)
	//log.Printf("Received a message: %s", str)
	data := strings.Split(str,",")
	var (
		id,_ = strconv.ParseInt(data[0],0,64)
		gameId,_ = strconv.ParseInt(data[1],10,0)
		//number = data[2]
	)
	var (
		gamePlayTemp []orm.Params
		//user model.Users
		//gameBetting model.GameBetting
		//balanceLog model.UserBalanceLogs
	)
	//游戏
	_,err := O.QueryTable(new(model.GamePlay)).Filter("id", id).Values(&gamePlayTemp,"id","number","status","is_status","prize_number","game_id")
	if err != nil{
		log.Printf("get gameplay fail, err: %s \n", err.Error())
		return
	}
	gamePlay := gamePlayTemp[0]
	if gamePlay["Status"] != 0{
		log.Printf("本期已经开奖，不用做其他处理")
		return
	}
	prizeNumber := gamePlay["PrizeNumber"]
	if gamePlay["IsStatus"] == 0 {
		//获取游戏的数据
		prizeNumber = getPrizeNumber(id, gameId)
	}
	fmt.Println(prizeNumber)
}

func myRedis()redis.Conn{
	databaseConf, err := config.NewConfig("ini","../../../conf/database.conf")
	if err != nil{
		fmt.Println("redis配置获取失败12")
		return nil
	}
	rConn,err := redis.Dial("tcp",databaseConf.String("redis_host") + ":" + databaseConf.String("redis_port"))
	if err != nil{
		fmt.Println("redis链接失败")
		return nil
	}
	if _,err = rConn.Do("AUTH",databaseConf.String("redis_pwd")); err != nil{
		fmt.Println(err)
		return nil
	}
	return rConn
}

func getPrizeNumber(gamePId, gameId int64)string{
	var data []orm.Params
	//用户投注数
	_,err := O.Raw(`SELECT sum(money) total_money FROM `+ model.DatabasePrefix +`game_betting WHERE game_p_id = ?`, gamePId).Values(&data,"total_money")
	if err != nil{
		fmt.Println(err)
		return ""
	}
	totalMoney := data[0]["total_money"]
	if totalMoney == nil{
		totalMoney = 0
	}
	RConn := myRedis()
	r,err := redis.String(RConn.Do("Get","laravel_database_GAME_CONFIG_" + strconv.FormatInt(gameId,10)))
	if err !=nil{
		fmt.Println(err)
		return ""
	}

	var gameConf model.Game
	e := json.Unmarshal([]byte(r), &gameConf)
	if e != nil{
		log.Println("json转化失败,err=>",e.Error())
		return ""
	}
	prizeNumber := ""
	openType := gameConf.OpenType.Value
	// 1天杀 2局杀 3随机
	switch openType {
	case 1:
		kill,_ := strconv.ParseFloat(gameConf.DateKill,10)
		prizeNumber = getDayKillNumber(gameId,gamePId,kill)
	case 2:
		kill,_ := strconv.ParseFloat(gameConf.OneKill,10)
		prizeNumber = getOneKillNumber(totalMoney.(float64),gamePId,gameId,kill)
	case 3:
		prizeNumber = GetRandNumberStr()
	}
	return prizeNumber
}

func GetRandNumberStr()string{
	randNum := rand.Intn(9)
	return strconv.FormatInt(int64(randNum),10)
}

func getDayKillNumber(gameId,id int64, kill float64)string{
	totalBetting,TotalWin := getAllDayBetting(gameId)
	canMoney := calcCanMoney(totalBetting, TotalWin, kill)
	return getNumber(canMoney,id,gameId)
}

func calcCanMoney(totalBetting, TotalWin, kill float64)float64{
	canMoney,_ := decimal.NewFromFloat(totalBetting).Sub(decimal.NewFromFloat(TotalWin)).Mul(decimal.NewFromFloat(1).Sub(decimal.NewFromFloat(kill))).Float64()
	return canMoney
}

func getOneKillNumber(totalBetting float64,id, gameId int64, kill float64)string{
	canMoney := calcCanMoney(totalBetting,0, kill)
	return getNumber(canMoney,id,gameId)
}

func getNumber(canMoney float64, id, gameId int64)string{
	return ""
}

func getAllDayBetting(gameId int64)(totalBetting, totalWinMoney float64){
	var data []orm.Params
	startTime,endTime := GetDayStartUnix()
	_,err := O.Raw("SELECT SUM(money) total_betting_money, SUM(win_money) total_win_money FROM "+ model.DatabasePrefix +"game_betting WHERE game_id = ? AND betting_time BETWEEN ? AND ?",gameId,startTime,endTime).Values(&data,"total_betting_money","total_win_money")
	if err != nil{
		fmt.Println(err)
		return 0,0
	}
	totalMoney := data[0]["total_betting_money"]
	if totalMoney == nil{
		totalMoney = 0
	}
	totalWin := data[0]["total_win_money"]
	if totalWin == nil{
		totalWin = 0
	}
	return totalMoney.(float64), totalWin.(float64)
}

func GetDayStartUnix()(start ,end int64){
	today := time.Now().Format(GetDayFormat())
	loc,_ := time.LoadLocation("Local") //获取时区
	tmpStart,_ := time.ParseInLocation(GetTimeFormat(),today + " 00:00:00",loc)
	startTime := tmpStart.Unix()
	tmpEnd,_ := time.ParseInLocation(GetTimeFormat(),today + " 23:59:59",loc)
	endTime := tmpEnd.Unix()
	return startTime, endTime
}

func GetTimeFormat()string{
	return "2006-01-02 15:04:05"
}

func GetDayFormat()string{
	return "2006-01-02"
}
