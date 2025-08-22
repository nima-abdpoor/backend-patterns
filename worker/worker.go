package worker

import (
	"backend-patterns/entity"
	"backend-patterns/outbox"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Worker struct {
	store       outbox.Store
	handler     func(task entity.Task) error
	pollEvery   time.Duration
	concurrency int
	stopCh      chan struct{}
	wg          sync.WaitGroup
}

func NewWorker(store outbox.Store, handler func(task entity.Task) error, pollEvery time.Duration, concurrency int) *Worker {
	return &Worker{
		store:       store,
		handler:     handler,
		pollEvery:   pollEvery,
		concurrency: concurrency,
		stopCh:      make(chan struct{}),
	}
}

func (w *Worker) Start() {
	for i := 0; i < w.concurrency; i++ {
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()
			ticker := time.NewTicker(w.pollEvery)
			defer ticker.Stop()
			for {
				select {
				case <-w.stopCh:
					return
				case <-ticker.C:
					w.tick()
				}
			}
		}()
	}
}

func (w *Worker) Stop() {
	close(w.stopCh)
	w.wg.Wait()
}

func (w *Worker) tick() {
	tasks, _ := w.store.Due(10)
	for _, t := range tasks {
		tt := t
		err := w.handler(tt)
		if err == nil {
			_ = w.store.MarkDone(tt.ID)
			continue
		}

		nextAttempts := tt.Attempts + 1
		if nextAttempts >= tt.MaxAttempts {
			_ = w.store.MarkFailed(tt.ID, err.Error())
			fmt.Println("worker ->", tt.ID, "err:", err)
			continue
		}

		backoff := computeBackoff(nextAttempts)
		nextTime := time.Now().Add(backoff)
		_ = w.store.Reschedule(tt.ID, nextTime, nextAttempts, err.Error())
	}
}

func computeBackoff(attempt int) time.Duration {
	base := 500 * time.Millisecond
	exp := float64(base) * math.Pow(2, float64(attempt-1))
	jitter := time.Duration(rand.Intn(250)) * time.Millisecond
	return time.Duration(exp) + jitter
}
