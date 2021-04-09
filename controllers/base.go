package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
	RtnJson
	O orm.Ormer
	Lang map[string]map[string]string
}

type RtnJson struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
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
