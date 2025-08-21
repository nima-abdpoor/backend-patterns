package order

import "fmt"

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s Service) Create() error {
	//todo we should create an order record in database with state: CREATED
	fmt.Println("order created.")
	return nil
}

func (s Service) CancelOrder() {
	// todo we can roll back the order record in our database
	//todo we can keep the order record and change the order state to CANCELED
	fmt.Println("order cancelled")
}

func (s Service) FinalizeOrder() error {
	//todo we should change the state of order record to COMPLETED
	fmt.Println("order cancelled")
	return nil
}
