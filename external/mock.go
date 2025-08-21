package external

import (
	"errors"
	"math/rand"
	"time"
)

type MockSms struct {
	timeout time.Duration
}

func NewMockSms(timeout time.Duration) SMS {
	return MockSms{
		timeout: timeout,
	}
}

func (s MockSms) SendSMS(to, msg string) error {
	r := rand.Intn(10)
	if r < 4 {
		time.Sleep(s.timeout)
		return errors.New("SMS API timeout")
	}
	if r < 6 {
		return errors.New("SMS API unavailable")
	}
	return nil
}
