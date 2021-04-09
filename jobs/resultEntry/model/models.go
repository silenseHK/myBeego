package model

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

var DatabasePrefix string

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	databaseConf, err := config.NewConfig("ini","../../conf/database.conf")
	if err != nil{
		fmt.Println("数据库配置获取失败,失败原因:" + err.Error())
		return
	}

	DatabasePrefix = databaseConf.String("prefix")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",
		databaseConf.String("user"),
		databaseConf.String("password"),
		databaseConf.String("host"),
		databaseConf.String("database"),
		databaseConf.String("charset"),
	)
	err2 := orm.RegisterDataBase("default", "mysql", dataSource)
	if err2 != nil{
		fmt.Println("数据库连接失败,失败愿因:" + err2.Error())
		return
	}

	var modelArr = []interface{}{
		new(GamePlay),
		new(Game),
	}
	orm.RegisterModelWithPrefix(DatabasePrefix, modelArr...)
}

type GamePlay struct{
	Id int `json:"id"`
	Number string `json:"number"`
	GameId int64 `json:"game_id"`
	PrizeNumber string `json:"prize_number"`
	Status int `json:"status"`
	PrizeTime int `json:"prize_time"`
	EndTime int64 `json:"end_time"`
	StartTime int64 `json:"start_time"`
	Type int `json:"type"`
	Winmoney float64 `json:"winmoney"`
	BMoney float64 `json:"b_money"`
	Lostmoney float64 `json:"lostmoney"`
	PtMoney float64 `json:"pt_money"`
	IsQueue int `json:"is_queue"`
	IsStatus int `json:"is_status"`
}

type Game struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
}
