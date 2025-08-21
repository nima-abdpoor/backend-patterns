package inventory

import "fmt"

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s Service) DeductInventory() error {
	//todo put lock in inventoryId in order to prevent multiple concurrent deductions example: we can use etcd or something simpler like redis...
	//todo deduct the count of InventoryId in database
	fmt.Println("Inventory deducted")
	return nil
}

func (s Service) RestoreInventory() {
	//todo put lock in inventoryId in order to prevent multiple concurrent restoration
	//todo increase the count of InventoryId or in other words update the counts to the value before DeductInventory
	fmt.Println("Inventory restored")
}
