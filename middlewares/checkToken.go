package middlewares

import (
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"hello/tools"
	"strconv"
	"strings"
)

func CheckUserToken (ctx *context.Context){
	header := ctx.Request.Header
	token,ok := header["Token"]
	if !ok {
		ctx.WriteString("缺少token")
		return
	}
	userId,err := tools.DecryptToken(fmt.Sprintf(strings.Join(token,"")))
	if err != nil{
		ctx.WriteString(err.Error())
		return
	}
	ctx.Input.SetParam("UserId",strconv.FormatInt(userId,10))
}

