package controller

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"hello/jobs/generateNumber/model"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

var O orm.Ormer

var timeRange int64 = 86400
//var timeRange int64 = 60 * 60

var gameIdx = map[int64]map[string]int64{}

var gameMap = map[int64]int64{
	4: 5 * 60,
	3: 3 * 60,
	2: 2 * 60,
	1: 1 * 60,
}

func init(){
	O = orm.NewOrm()
}

func GenerateNumber(){
	var game []model.Game
	_,err := O.QueryTable(new(model.Game)).OrderBy("id").All(&game)
	if err != nil{
		fmt.Println(err)
		return
	}
	var gameChan = make(chan model.GamePlay,3000)
	var flag = make(chan int64,4)
	for _,v := range game{
		wg.Add(1)
		go perGameAct(gameChan, v.Id, flag)
	}
	wg.Add(1)
	go insertGamePlay(gameChan)
	wg.Wait()
	fmt.Println("生成完成")
}

func perGameAct(gameChan chan model.GamePlay,gameId int64, flag chan int64){
	defer wg.Done()
	var preGamePlay model.GamePlay
	e := O.QueryTable(new(model.GamePlay)).Filter("game_id",gameId).OrderBy("-end_time").One(&preGamePlay)
	if e != nil{
		return
	}
	startTime := preGamePlay.EndTime
	for i:=timeRange/gameMap[gameId];i>0;i--{
		makeGamePlay(gameChan, gameId, &startTime)
	}
	flag <- 1
	if len(flag) == 4{
		close(gameChan)
	}
	return
}

func makeNumber(gameId, startTime int64)string{
	if gameIdx[gameId] == nil{
		gameIdx[gameId] = map[string]int64{}
	}
	idx := gameIdx[gameId][time.Unix(startTime, 0).Format("2006010215")]
	if idx == 0{
		gameIdx[gameId][time.Unix(startTime, 0).Format("2006010215")] = 1
	}else{
		gameIdx[gameId][time.Unix(startTime, 0).Format("2006010215")] += 1
	}
	idx = gameIdx[gameId][time.Unix(startTime, 0).Format("2006010215")]
	gameIdxStr := strconv.FormatInt(idx,10)
	for i:=5-len(gameIdxStr);i>0;i--{
		gameIdxStr = "0" + gameIdxStr
	}
	return time.Unix(startTime, 0).Format("2006010215") + strconv.FormatInt(gameId,10) + gameIdxStr
}

func makeGamePlay(gameChan chan model.GamePlay, gameId int64, startTime *int64){
	var gamePlay = model.GamePlay{
		GameId: gameId,
		Number: makeNumber(gameId, *startTime),
		StartTime: *startTime,
		EndTime: *startTime + gameMap[gameId],
	}
	*startTime = gamePlay.EndTime
	gameChan <- gamePlay
}

func insertGamePlay(gameChan chan model.GamePlay){
	defer wg.Done()
	for v := range gameChan{
		O.Insert(&v)
	}
}