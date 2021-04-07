package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

const KEY = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()"

func EncryptToken(userId int64)string{
	dataString := strconv.FormatInt(userId,10) + ":" + strconv.FormatInt(time.Now().Unix(),10)
	data := jwt.StandardClaims{Subject: dataString, ExpiresAt: time.Now().Unix() + int64(time.Hour * 2)}
	tokenInfo := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenStr,err := tokenInfo.SignedString([]byte(KEY))
	if err != nil{
		fmt.Println(err)
	}
	return tokenStr
}

func DecryptToken(tokenStr string)(userId int64, err error){
	tokenInfo,_ := jwt.Parse(tokenStr, func(token *jwt.Token)(i interface{}, e error){
		return KEY,nil
	})
	e := tokenInfo.Claims.Valid()
	if e != nil{
		return 0,e
	}
	finToken := tokenInfo.Claims.(jwt.MapClaims)
	ok := finToken.VerifyExpiresAt(time.Now().Unix(),true)
	if !ok {
		return 0,errors.New("token 过期")
	}
	str := finToken["sub"].(string)
	byteData := strings.Split(str,":")
	userId,err = strconv.ParseInt(byteData[0],10,64)
	return
}
