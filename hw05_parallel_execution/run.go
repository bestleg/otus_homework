package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errLimit")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n int, errLimit int) error {
	taskCh := make(chan Task, len(tasks))
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}

	for _, t := range tasks {
		taskCh <- t
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(taskCh, &errLimit, &mu)
		}()
	}

	close(taskCh)
	wg.Wait()

	if errLimit <= 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func worker(tasksCh <-chan Task, errLimit *int, mu *sync.RWMutex) {
	stop := false
	for task := range tasksCh {
		mu.RLock()
		if *errLimit < 1 {
			stop = true
		}
		mu.RUnlock()
		if stop {
			return
		}
		err := task()
		mu.Lock()
		if err != nil {
			*errLimit--
		}
		mu.Unlock()
	}
}
