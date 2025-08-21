package checkout

import (
	"backend-patterns/services/inventory"
	"backend-patterns/services/order"
	"backend-patterns/services/payment"
	"fmt"
)

type Service struct {
	orderSvc     order.Service
	inventorySvc inventory.Service
	paymentSvc   payment.Service
}

func NewService(order order.Service, inventory inventory.Service, payment payment.Service) Service {
	return Service{
		orderSvc:     order,
		inventorySvc: inventory,
		paymentSvc:   payment,
	}
}

// Checkout Saga Orchestrator
func (s Service) Checkout() error {
	if err := s.orderSvc.Create(); err != nil {
		return err
	}

	if err := s.inventorySvc.DeductInventory(); err != nil {
		s.orderSvc.CancelOrder()
		return err
	}

	if err := s.paymentSvc.ProcessPayment(); err != nil {
		s.inventorySvc.RestoreInventory()
		s.orderSvc.CancelOrder()
		return err
	}

	if err := s.orderSvc.FinalizeOrder(); err != nil {
		s.inventorySvc.RestoreInventory()
		s.paymentSvc.RefundPayment()
		s.orderSvc.CancelOrder()
		return err
	}

	fmt.Println("Checkout completed successfully")
	return nil
}
