package outbox

import (
	"backend-patterns/entity"
	"errors"
	"fmt"
	"sync"
	"time"
)

type InMemoryOutbox struct {
	mu    sync.Mutex
	tasks map[string]*entity.Task
}

func NewInMemoryOutbox() *InMemoryOutbox {
	return &InMemoryOutbox{tasks: make(map[string]*entity.Task)}
}

func (s *InMemoryOutbox) Enqueue(t entity.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	tt := t
	fmt.Printf("publishing new task with Id: %s into queue...\n", t.ID)
	s.tasks[t.ID] = &tt
	return nil
}

func (s *InMemoryOutbox) Due(limit int) ([]entity.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	out := make([]entity.Task, 0, limit)
	for _, t := range s.tasks {
		if t.Status == entity.StatusPending && !t.NextAttempt.After(now) {
			out = append(out, *t)
			if len(out) == limit {
				break
			}
		}
	}
	return out, nil
}

func (s *InMemoryOutbox) MarkDone(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("marking a task as done. taskId: %s\n", id)
	if t, ok := s.tasks[id]; ok {
		t.Status = entity.StatusDone
		return nil
	}
	return errors.New("task not found")
}

func (s *InMemoryOutbox) MarkFailed(id string, err string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if t, ok := s.tasks[id]; ok {
		t.Status = entity.StatusFailed
		t.LastError = err
		return nil
	}
	return errors.New("task not found")
}

func (s *InMemoryOutbox) Reschedule(id string, next time.Time, attempts int, lastErr string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if t, ok := s.tasks[id]; ok {
		t.Attempts = attempts
		t.NextAttempt = next
		t.LastError = lastErr
		return nil
	}
	return errors.New("task not found")
}
