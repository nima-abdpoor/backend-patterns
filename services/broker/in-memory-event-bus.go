package broker

import "sync"

type InMemoryEventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]chan Event
}

func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		subscribers: make(map[string][]chan Event),
	}
}

func (eb *InMemoryEventBus) Subscribe(eventName string, ch chan Event) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.subscribers[eventName] = append(eb.subscribers[eventName], ch)
}

func (eb *InMemoryEventBus) Publish(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	if chans, ok := eb.subscribers[event.Name]; ok {
		for _, ch := range chans {
			go func(c chan Event) {
				c <- event
			}(ch)
		}
	}
}
