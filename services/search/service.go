package search

import (
	"backend-patterns/services/broker"
	"fmt"
)

type Service struct{}

func NewService(b broker.EventBus) *Service {
	s := &Service{}
	ch := make(chan broker.Event)

	b.Subscribe(broker.EventCourseCreated, ch)

	go s.IndexSearch(ch)

	return s
}

func (s *Service) IndexSearch(ch chan broker.Event) {
	for e := range ch {
		fmt.Println("Indexing course in search system:", e.Data)
	}
}
