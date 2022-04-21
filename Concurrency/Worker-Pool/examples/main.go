package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func main() {
	numJobs := 5
	job := make(chan int, numJobs)
	result := make(chan int, numJobs)
	for w := 1; w <= 3; w++ {
		go worker(w, job, result)
	}
	for j := 1; j <= numJobs; j++ {
		job <- j
	}
	close(job)
	for a := 1; a <= numJobs; a++ {
		<-result
	}
}
