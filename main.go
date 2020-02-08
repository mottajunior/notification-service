package main

import (
	"fmt"
	rabbitMQ "github.com/mottajunior/notification-service/consumer"

)

func main(){
	fmt.Println("Notification Service has started")
	rabbitMQ.ConsumeQueue()
}


