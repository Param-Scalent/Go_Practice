package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go task1(&wg)
	go task2(&wg)
	go task3(&wg)
	go task4(&wg)
	go task5(&wg)
	wg.Wait()

}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	// http request
	_, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("could not make http request: %v", err)
	}
	fmt.Println("task 1 is done")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	// crazy calculations
	var count int
	for i := 0; i < 1_000_000_000; i++ {
		count += i
	}
	fmt.Println("task 2 is done")
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task 3 is done")
}

func task4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Microsecond)
	fmt.Println("task 4 is done")
}

func task5(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(600 * time.Microsecond)
	fmt.Println("task 5 is done")
}
