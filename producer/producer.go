package producer

import (
	"backend-patterns/entity"
	"backend-patterns/outbox"
	"encoding/json"
	"time"
)

func EnqueueUserProfileUpdate(store outbox.Store, evt entity.UserProfileUpdate) error {
	payload, _ := json.Marshal(evt)
	t := entity.Task{
		ID:          "task-" + evt.UserID + "", //todo this should be unique number like ULID
		Type:        "UserProfileUpdated",
		Payload:     payload,
		Attempts:    0,
		MaxAttempts: 6,
		NextAttempt: time.Now(),
		Status:      entity.StatusPending,
		CreatedAt:   time.Now(),
	}
	return store.Enqueue(t)
}
