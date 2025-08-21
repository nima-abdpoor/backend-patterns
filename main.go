package main

import (
	"backend-patterns/services/checkout"
	"backend-patterns/services/inventory"
	"backend-patterns/services/order"
	"backend-patterns/services/payment"
	"fmt"
)

func main() {
	orderSvc := order.NewService()
	inventorySvc := inventory.NewService()
	paymentSvc := payment.NewService()

	checkoutSvc := checkout.NewService(orderSvc, inventorySvc, paymentSvc)

	if err := checkoutSvc.Checkout(); err != nil {
		fmt.Printf("checkout process failed with error: %v\n", err.Error())
	} else {
		fmt.Println("resilient checkout process successfully done")
	}
}
