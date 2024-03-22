package main

import (
	"fmt"
	"sync"
	"time"
)

type MyFunction func(x int)

type RateLimiter struct {
	limit          time.Duration
	initialWindow  time.Duration
	taskCount      int
	mutex          sync.Mutex
	resetThreshold int
}

func NewRateLimiter(limit, initialWindow time.Duration, resetThreshold int) *RateLimiter {
	return &RateLimiter{
		limit:          limit,
		initialWindow:  initialWindow,
		resetThreshold: resetThreshold,
	}
}

func (r *RateLimiter) Execute(task MyFunction, param int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Reset task count if elapsed time exceeds initial window
	if time.Since(startTime) >= r.initialWindow {
		r.taskCount = 0
		startTime = time.Now()
	}

	// If within initial window and task count is below threshold, execute immediately
	if r.taskCount < r.resetThreshold {
		r.taskCount++
	} else {
		// If outside initial window or task count exceeds threshold, rate-limit
		elapsed := time.Since(startTime)
		if elapsed < r.initialWindow {
			time.Sleep(r.initialWindow - elapsed)
		}
		time.Sleep(r.limit)
	}

	// Execute task
	task(param)
}

var startTime time.Time

func main() {
	// Initialize a rate limiter with a limit of 1 call per second and 5 tasks allowed to run immediately in the first 5 seconds
	limiter := NewRateLimiter(time.Second, 10*time.Second, 10)

	startTime = time.Now()
	fmt.Printf("Task started at: %02d:%02d\n", startTime.Minute(), startTime.Second())

	var caller MyFunction = func(num int) {
		fmt.Println(fmt.Sprintf("Executing task %d ...", num))
		startTime := time.Now()
		fmt.Printf("Task started at: %02d:%02d\n", startTime.Minute(), startTime.Second())
		fmt.Println(fmt.Sprintf("Task %d completed.", num))
	}
	// Simulate 20 calls to Execute, which should be rate-limited after the initial window
	for i := 0; i < 20; i++ {
		limiter.Execute(caller, i)
	}

	stopTime := time.Now()
	fmt.Printf("Task finished at: %02d:%02d\n", stopTime.Minute(), stopTime.Second())
}
