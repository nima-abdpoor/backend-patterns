package external

type SMS interface {
	SendSMS(to, msg string) error
}
