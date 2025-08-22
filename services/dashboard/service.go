package dashboard

import (
	"backend-patterns/services/broker"
	"fmt"
)

type Service struct{}

func NewService(b broker.EventBus) *Service {
	s := &Service{}
	ch := make(chan broker.Event)

	b.Subscribe(broker.EventCourseCreated, ch)
	go s.UpdateDashboard(ch)

	return s
}

func (s *Service) UpdateDashboard(ch chan broker.Event) {
	for e := range ch {
		fmt.Println("Updating dashboard with course:", e.Data)
	}
}
