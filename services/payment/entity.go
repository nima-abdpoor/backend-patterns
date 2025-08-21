package payment

type State string

const (
	Created   State = "CREATED"
	Failure   State = "FAILURE"
	Success   State = "SUCCESS"
	Refunding State = "REFUNDING"
	Refunded  State = "REFUNDED"
)

type Payment struct {
	PaymentId          string
	OrderId            string
	SourceAccount      string
	DestinationAccount string
	Amount             int64
	State              State
}
