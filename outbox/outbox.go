package outbox

import (
	"backend-patterns/entity"
	"time"
)

type Store interface {
	Enqueue(t entity.Task) error
	Due(limit int) ([]entity.Task, error)
	MarkDone(id string) error
	MarkFailed(id string, err string) error
	Reschedule(id string, next time.Time, attempts int, lastErr string) error
}
