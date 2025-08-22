package main

import (
	"backend-patterns/services/broker"
	"backend-patterns/services/dashboard"
	"backend-patterns/services/notification"
	"backend-patterns/services/search"
	"time"
)

func main() {
	eventBus := broker.NewInMemoryEventBus()

	dashboard.NewService(eventBus)
	notification.NewService(eventBus)
	search.NewService(eventBus)

	eventBus.Publish(broker.Event{
		Name: broker.EventCourseCreated,
		Data: "Scrum In 10 Hours!!",
	})

	time.Sleep(1 * time.Second)
}
