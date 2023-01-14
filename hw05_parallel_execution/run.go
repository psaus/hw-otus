package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrInvalidWorkerNumber = errors.New("invalid worker number")
)

type Task func() error

func worker(wg *sync.WaitGroup, taskChan <-chan Task, errorCounter *SafeCounter) {
	defer wg.Done()

	for t := range taskChan {
		if err := t(); err != nil {
			errorCounter.Increment()
		}
	}
}

type SafeCounter struct {
	sync.Mutex
	count int
}

func (s *SafeCounter) Increment() {
	s.Mutex.Lock()
	s.count++
	s.Mutex.Unlock()
}

func (s *SafeCounter) Value() int {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.count
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return nil
	}

	if 0 >= n {
		return ErrInvalidWorkerNumber
	}

	wg := &sync.WaitGroup{}
	taskChan := make(chan Task)
	errorCounter := &SafeCounter{}

	// spawning workers
	wg.Add(1)
	go func() {
		defer wg.Done()

		if len(tasks) < n {
			n = len(tasks)
		}

		for i := 0; i < n; i++ {
			wg.Add(1)
			go worker(wg, taskChan, errorCounter)
		}
	}()

	var wasErrors bool
	// channel producer
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(taskChan)

		for _, task := range tasks {
			wasErrors = m > 0 && errorCounter.Value() >= m
			if wasErrors {
				break
			}

			taskChan <- task
		}
	}()

	wg.Wait()

	if wasErrors {
		return ErrErrorsLimitExceeded
	}

	return nil
}
