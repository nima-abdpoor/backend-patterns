package service

import (
	"backend-patterns/entity"
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

type Service struct {
	mu       sync.Mutex
	applied  map[string]bool
	storage  map[string]entity.UserProfileUpdate
	failRate int
}

func NewProfileService(failRate int) *Service {
	return &Service{
		applied:  make(map[string]bool),
		storage:  make(map[string]entity.UserProfileUpdate),
		failRate: failRate,
	}
}

func (p *Service) ApplyUpdate(evt entity.UserProfileUpdate) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	//we shouldn't apply a single update twice. so we simulate the idempotency in this way...
	if p.applied[evt.IdemKey] {
		return nil
	}

	if rand.Intn(10) < p.failRate {
		fmt.Println("failed to update user profile due to an error...")
		return errors.New("temporary downstream error")
	} else {
		fmt.Println("user profile updates applied successfully")
	}

	p.applied[evt.IdemKey] = true
	p.storage[evt.UserID] = evt
	return nil
}
