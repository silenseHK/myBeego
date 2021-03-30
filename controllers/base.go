package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	RtnJson
	O orm.Ormer
}

const (
	Success int = 200
	Fail int = 300
)

type RtnJson struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *BaseController) Prepare(){
	c.RtnJson.Code = Success
	c.RtnJson.Msg = "success"
	c.RtnJson.Data = ""

	c.O = orm.NewOrm()
}

func (c *BaseController) ReturnJson(){
	rtn, _ := json.Marshal(c.RtnJson)
	c.Ctx.WriteString(string(rtn))
}

func (c *BaseController) ReturnJsonWithData(code int, msg string, data interface{}){
	c.RtnJson.Code = code
	c.RtnJson.Msg = msg
	c.RtnJson.Data = data
	c.ReturnJson()
}
