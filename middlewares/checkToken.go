package middlewares

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"hello/tools"
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

	//newToken := tools.EncryptToken(12345678)
	//fmt.Println(newToken)
	userId,err := tools.DecryptToken(fmt.Sprintf(strings.Join(token,"")))
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(userId)
	fmt.Println(ctx.Request.Body)
	return
}

