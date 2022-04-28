package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementor int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)
			fmt.Println("Counter mid val: ", atomic.LoadInt64(&incrementor))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter End val: ", incrementor)
}
