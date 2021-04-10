package main

import (
	"fmt"
	"hello/jobs/rabbitMq/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("result_game" +
		"silence")

	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello silence!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
