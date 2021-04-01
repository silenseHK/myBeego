package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"hello/libs"
	"net/http"
)

type Middleware func(handler http.Handler) http.Handler

type BaseController struct {
	beego.Controller
	RtnJson
	O orm.Ormer
}

type RtnJson struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *BaseController) Prepare(){
	c.RtnJson.Code = libs.Success
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

//func (c *BaseController) Validate(params map[string]string)map[string]interface{}{
//	var rtn map[string]interface{}
//	for k,v := range params{
//		c.perValidate(&rtn,k,v)
//	}
//	return rtn
//}
//
//func (c *BaseController) perValidate(rtn *map[string]interface{}, field string, kind string){
//	kind = "Get" + strings.ToUpper(string([]byte(kind)[0])) + string([]byte(kind)[1:len(kind)])
//	reflect.ValueOf(c)
//}
