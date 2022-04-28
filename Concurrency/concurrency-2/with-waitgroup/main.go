package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	now := time.Now()
	for i := 0; i < 10; i++ {
		go work(&wg, i+1)
	}

	wg.Wait()
	fmt.Println("elapsed", time.Since(now))
	fmt.Println("main is done")
}

func work(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task", id, "is done")

}
