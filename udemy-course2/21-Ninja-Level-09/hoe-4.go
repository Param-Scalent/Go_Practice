package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incrementor := 0
	gs := 100

	var m sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementor
			v++
			incrementor = v
			fmt.Println("Counter mid val: ", incrementor)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter End val: ", incrementor)
}
