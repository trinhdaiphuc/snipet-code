package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 20)
	results := make(chan int, 20)
	numWorkers := 4
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < 20; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 20; j++ {
		<-results
	}
}
