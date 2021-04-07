package tools

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/garyburd/redigo/redis"
)

func init(){
	databaseConf, err := config.NewConfig("ini","conf/database.conf")
	if err != nil{
		fmt.Println("redis配置获取失败")
		return
	}
	conn,err := redis.Dial("tcp",databaseConf.String("redis_host") + ":" + databaseConf.String("redis_port"))
	if err != nil{
		fmt.Println("redis链接失败")
		return
	}
	defer conn.Close()
	if _,err = conn.Do("AUTH",databaseConf.String("redis_pwd")); err != nil{
		fmt.Println(err)
		return
	}

}
