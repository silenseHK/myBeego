package main

import (
	"hello/jobs/generateNumber/controller"
	_ "hello/jobs/generateNumber/model"
)

func init(){

}

func main(){
	controller.GenerateNumber()
}
