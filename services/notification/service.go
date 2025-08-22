package notification

import (
	"backend-patterns/services/broker"
	"fmt"
)

type Service struct{}

func NewService(b broker.EventBus) *Service {
	s := &Service{}
	ch := make(chan broker.Event)

	b.Subscribe(broker.EventCourseCreated, ch)

	go s.SendEmail(ch)

	return s
}

func (s *Service) SendEmail(ch chan broker.Event) {
	for e := range ch {
		fmt.Println("Sending email for new course:", e.Data)
	}
}
