package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker ", id, " started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker ", id, " started job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 1; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		fmt.Println("Sending a job to start ", j)
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println("Starting a jobs ", a)
		<-results
	}

}
