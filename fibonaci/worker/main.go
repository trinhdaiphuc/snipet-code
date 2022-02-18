package main

import (
	"fmt"
	"time"
)

func main() {
	defer elapsed()()
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	workerNums := 4

	// concurrent workers
	// remove or add some to test different configs
	for i := 0; i < workerNums; i++ {
		go worker(jobs, results)
	}

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 50; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// calculate time elapsed
func elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Calculation took %v\n", time.Since(start))
	}
}
