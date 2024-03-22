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
	waitCh         chan struct{} // Channel to control the number of waiting goroutines
}

func NewRateLimiter(limit, initialWindow time.Duration, resetThreshold, maxWaitingGoroutines int) *RateLimiter {
	// Initialize the channel with the specified buffer size and fill it to capacity
	waitCh := make(chan struct{}, maxWaitingGoroutines)
	for i := 0; i < maxWaitingGoroutines; i++ {
		waitCh <- struct{}{}
	}
	return &RateLimiter{
		limit:          limit,
		initialWindow:  initialWindow,
		resetThreshold: resetThreshold,
		waitCh:         waitCh,
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
	}

	// Retrieve token from the channel, effectively blocking if buffer is empty
	<-r.waitCh

	// Execute task
	task(param)

	// Release token back to the channel
	r.waitCh <- struct{}{}
}

var startTime time.Time

func main() {
	// Initialize a rate limiter with a limit of 1 call per second and 5 tasks allowed to run immediately in the first 5 seconds
	limiter := NewRateLimiter(time.Second, 10*time.Second, 10, 2)

	startTime = time.Now()
	fmt.Printf("Task started at: %02d:%02d\n", startTime.Minute(), startTime.Second())

	var caller MyFunction = func(num int) {
		fmt.Println(fmt.Sprintf("Executing task %d ...", num))
		startTime := time.Now()
		fmt.Printf("Task started at: %02d:%02d\n", startTime.Minute(), startTime.Second())
		time.Sleep(500 * time.Millisecond) // Simulating some work
		fmt.Println(fmt.Sprintf("Task %d completed.", num))
	}

	// Simulate 20 calls to Execute, which should be rate-limited after the initial window
	for i := 0; i < 20; i++ {
		limiter.Execute(caller, i)
	}

	stopTime := time.Now()
	fmt.Printf("Task finished at: %02d:%02d\n", stopTime.Minute(), stopTime.Second())
}
