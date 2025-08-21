package circuitbreaker

import (
	"sync"
	"time"
)

const (
	Closed = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	mu            sync.Mutex
	state         int
	failures      int
	successes     int
	failureThresh int
	resetTimeout  time.Duration
	lastFailure   time.Time
}

func NewCircuitBreaker(failureThresh int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:         Closed,
		failureThresh: failureThresh,
		resetTimeout:  resetTimeout,
	}
}

func (cb *CircuitBreaker) Allow() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case Open:
		if time.Since(cb.lastFailure) > cb.resetTimeout {
			cb.state = HalfOpen
			return true
		}
		return false
	case HalfOpen, Closed:
		return true
	default:
		return true
	}
}

func (cb *CircuitBreaker) Success() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.state == HalfOpen {
		cb.successes++
		if cb.successes > 2 {
			cb.state = Closed
			cb.failures = 0
			cb.successes = 0
		}
	} else if cb.state == Closed {
		cb.failures = 0
	}
}

func (cb *CircuitBreaker) Failure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.failures++
	cb.lastFailure = time.Now()

	if cb.state == HalfOpen || (cb.state == Closed && cb.failures >= cb.failureThresh) {
		cb.state = Open
	}
}
