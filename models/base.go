package models

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

var DatabasePrefix string

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	databaseConf, err := config.NewConfig("ini","conf/database.conf")
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
		new(Users),
		new(GameBetting),
		new(GameConfig),
		new(UserBalanceLogs),
	}

	orm.RegisterModelWithPrefix(DatabasePrefix, modelArr...)
}

