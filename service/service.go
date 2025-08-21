package service

import (
	"backend-patterns/circuitbreaker"
	"backend-patterns/external"
	"fmt"
)

type Service struct {
	sms     external.SMS
	breaker *circuitbreaker.CircuitBreaker
}

func NewService(sms external.SMS, cb *circuitbreaker.CircuitBreaker) Service {
	return Service{
		sms:     sms,
		breaker: cb,
	}
}

func (s Service) SendMessage(req SendSMSRequest) (SendSMSResponse, error) {
	var rsp SendSMSResponse

	if !s.breaker.Allow() {
		//todo we can create task for this sms and retry it later...
		//todo this should be a log
		//todo we can put metric here to count how many lost sms we have due to unavailable service
		fmt.Println("Circuit OPEN. service is not available")
		return rsp, fmt.Errorf("SMS API is not available right now, we will retry")
	}

	err := s.sms.SendSMS(req.To, req.Message)
	if err != nil {
		//todo this should be a log
		//we can retry if needed
		fmt.Println("SMS failed:", err)
		s.breaker.Failure()
		return rsp, fmt.Errorf("failed to send sms")
	}

	//todo this should be a log
	fmt.Println("SMS sent successfully")
	s.breaker.Success()

	rsp = SendSMSResponse{
		Success: true,
	}

	return rsp, nil
}
