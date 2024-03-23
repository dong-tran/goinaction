package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for {
		j, ok := <-jobs
		if ok {
			fmt.Println("worker", id, "started  job", j)
			// time.Sleep(time.Second)
			fmt.Println("worker", id, "finished job", j)
			results <- j * 2
		}
	}
}

func main() {
	var wg sync.WaitGroup
	const numJobs = 1000
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			go worker(x, jobs, results)
		}(w)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		s, ok := <-results
		if ok {
			fmt.Printf("Result: %d\n", s)
		}
	}
	wg.Wait()
}
