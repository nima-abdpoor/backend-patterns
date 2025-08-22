package main

import (
	"backend-patterns/entity"
	"backend-patterns/outbox"
	"backend-patterns/producer"
	"backend-patterns/service"
	worker2 "backend-patterns/worker"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//we are using outbox pattern
	store := outbox.NewInMemoryOutbox()

	//mocking the fail rate to make service unavailable
	profileSvc := service.NewProfileService(6)

	//we will use task queue to create task for retry handling and distribute tasks over workers
	worker := worker2.NewWorker(store, func(t entity.Task) error {
		switch t.Type {
		case "UserProfileUpdated":
			var evt entity.UserProfileUpdate
			if err := json.Unmarshal(t.Payload, &evt); err != nil {
				return err
			}
			return profileSvc.ApplyUpdate(evt)
		default:
			return fmt.Errorf("unknown task type: %s", t.Type)
		}
	}, 300*time.Millisecond, 2)
	worker.Start()
	defer worker.Stop()

	// create new task and profile update
	update := entity.UserProfileUpdate{
		UserID: "u-123",
		Name:   "Nima Abdpoor",
		Email:  "nimaabdpoor@gmail.com",
	}
	_ = producer.EnqueueUserProfileUpdate(store, update)

	time.Sleep(6 * time.Second)

	fmt.Println("task done")
}
