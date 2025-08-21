package payment

import (
	"errors"
	"fmt"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s Service) ProcessPayment() error {
	//todo create payment record in database with state like CREATED

	//todo call external service for process payment
	//todo if error happened update the record and change the state to FAILURE
	return errors.New("payment gateway down")

	//todo if it was successful update the record and change it to SUCCESS

	//fmt.Println("Payment processed.")
	//return nil
}

func (s Service) RefundPayment() {
	//todo we should change the record and state to REFUNDING
	//todo Call refunding external service
	//todo if refunding failed, we can create task or job (in our distributed task queue manager) to retry it later or simply run a job to retry all those records with REFUNDING state.
	//todo if refunding succeed, we should change the state to REFUNDED
	fmt.Println("Payment refunded")
}
