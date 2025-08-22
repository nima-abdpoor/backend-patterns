package broker

type Event struct {
	Name string
	Data interface{}
}

type EventBus interface {
	Subscribe(eventName string, ch chan Event)
	Publish(event Event)
}

const EventCourseCreated = "COURSE_CREATED"
