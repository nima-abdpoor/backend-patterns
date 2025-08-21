package main

import (
	"backend-patterns/circuitbreaker"
	"backend-patterns/external"
	"backend-patterns/service"
	"fmt"
	"time"
)

func main() {
	cb := circuitbreaker.NewCircuitBreaker(3, 5*time.Second)
	smsAPI := external.NewMockSms(10 * time.Second)
	myService := service.NewService(smsAPI, cb)

	//simulate the service calling

	for i := 0; i < 15; i++ {
		result, err := myService.SendMessage(service.SendSMSRequest{
			To:      "09123456789",
			Message: "Hello!",
		})

		if err != nil {
			fmt.Printf("can not send sms due to:%v\n", err.Error())
			//todo return proper error to user
		} else {
			fmt.Printf("sms sent successfully, success:%v\n", result.Success)
		}
	}
}
