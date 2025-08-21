package order

type State string

const (
	Created   State = "CREATED"
	Canceled  State = "CANCELED"
	Completed State = "COMPLETED"
)

type Order struct {
	ID           string
	UserID       string
	Status       string
	InventoryIds []string
}
