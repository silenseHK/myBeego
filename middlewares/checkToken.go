package middlewares

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"hello/tools"
	"strconv"
	"strings"
)

func CheckUserToken(ctx *context.Context){
	header := ctx.Request.Header
	fmt.Println(header)
	token,ok := header["Token"]
	if !ok {
		ctx.WriteString("缺少token")
		return
	}

	userId,err := tools.DecryptToken(fmt.Sprintf(strings.Join(token,"")))
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(userId)
	ctx.Input.SetParam("UserId",strconv.FormatInt(userId,10))
	return
}

