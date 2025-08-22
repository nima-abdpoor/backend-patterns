package entity

import "time"

type Task struct {
	ID          string
	Type        string
	Payload     []byte
	Attempts    int
	MaxAttempts int
	NextAttempt time.Time
	Status      TaskStatus
	CreatedAt   time.Time
	LastError   string
}

type TaskStatus string

const (
	StatusPending TaskStatus = "PENDING"
	StatusDone    TaskStatus = "DONE"
	StatusFailed  TaskStatus = "FAILED"
)
