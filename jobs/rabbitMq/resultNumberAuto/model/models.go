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
	databaseConf, err := config.NewConfig("ini","../../../conf/database.conf")
	if err != nil{
		fmt.Println("数据库配置获取失败,失败原因1:" + err.Error())
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
		new(Users),
		new(UserBalanceLogs),
		new(GameBetting),
	}
	orm.RegisterModelWithPrefix(DatabasePrefix, modelArr...)
}

type GamePlay struct{
	Id int `json:"id" mysql:"id"`
	Number string `json:"number" mysql:"number"`
	GameId int64 `json:"game_id" mysql:"game_id"`
	PrizeNumber string `json:"prize_number" mysql:"prize_number"`
	Status int `json:"status" mysql:"status"`
	PrizeTime int `json:"prize_time" mysql:"prize_time"`
	EndTime int64 `json:"end_time" mysql:"end_time"`
	StartTime int64 `json:"start_time" mysql:"start_time"`
	Type int `json:"type" mysql:"type"`
	Winmoney float64 `json:"winmoney" mysql:"winmoney"`
	BMoney float64 `json:"b_money" mysql:"b_money"`
	Lostmoney float64 `json:"lostmoney" mysql:"lostmoney"`
	PtMoney float64 `json:"pt_money" mysql:"pt_money"`
	IsQueue int `json:"is_queue" mysql:"is_queue"`
	IsStatus int `json:"is_status" mysql:"is_status"`
}

type UserBalanceLogs struct{
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Type int8 `json:"type"`
	DqBalance float64 `json:"dq_balance"`
	WcBalance float64 `json:"wc_balance"`
	Time int64 `json:"time"`
	Msg string `json:"msg"`
	Money float64 `json:"money"`
	IsFirstRecharge int8 `json:"is_first_recharge"`
	AdminId int `json:"admin_id"`
}

type Users struct{
	Id int `json:"id"`
	Phone string `json:"phone"`
	Balance float64 `json:"balance"`
}

type GameBetting struct {
	Id             int `json:"id"`
	BettingNum     int64 `json:"betting_num"`
	UserId         int `json:"user_id"`
	GameId         int `json:"game_id"`
	GamePId        int `json:"game_p_id"`
	GameCXId       int `json:"game_c_x_id"`
	Money          float64 `json:"money"`
	Odds           float64 `json:"odds"`
	WinMoney       int `json:"win_money"`
	BettingTime    int64 `json:"betting_time"`
	SettlementTime int `json:"settlement_time"`
	Status         int `json:"status"`
	Type           int `json:"type"`
	ServiceCharge  float64 `json:"service_charge"`
}

type Game struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Weight int `json:"weight"`
	Status string `json:"status"`
	OpenType OpenType `json:"open_type"`
	DateKill string `json:"date_kill"`
	OneKill string `json:"one_kill"`
	LockTime int64 `json:"lock_time"`
}

type OpenType struct{
	Value int `json:"value"`
	Title string `json:"title"`
}