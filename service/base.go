package service

import (
	"github.com/beego/beego/v2/client/orm"
)

var O orm.Ormer

type BaseService struct{

}

func init(){
	O = orm.NewOrm()
}
