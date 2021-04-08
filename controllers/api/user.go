package api

import (
	"hello/libs"
	"hello/tools"
)

type UserController struct {
	BaseApiController
}

func (c UserController)Login(){
	//phone := c.GetString("phone")
	//if len(phone) != 8{
	//	c.ReturnJsonWithData(libs.Fail,"请输入正确的登陆账号","")
	//	return
	//}
	c.ReturnJsonWithData(libs.Success,"",tools.EncryptToken(256))
	return
}