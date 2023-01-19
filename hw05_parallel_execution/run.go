package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrInvalidWorkerNumber = errors.New("invalid worker number")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return nil
	}

	if 0 >= n {
		return ErrInvalidWorkerNumber
	}

	taskChan := make(chan Task)

	wg := &sync.WaitGroup{}
	wg.Add(n)
	var errorCounter int32

	// spawning workers
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			// read the channel until it closes
			for t := range taskChan {
				if err := t(); err != nil {
					atomic.AddInt32(&errorCounter, 1)
				}
			}
		}()
	}

	// channel producer
	for _, task := range tasks {
		if m > 0 && int(atomic.LoadInt32(&errorCounter)) >= m {
			close(taskChan)
			wg.Wait()
			return ErrErrorsLimitExceeded
		}

		taskChan <- task
	}

	close(taskChan)
	wg.Wait()
	return nil
}
