package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	//send
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from eve Channel: ", v)
		case v := <-o:
			fmt.Println("from odd Channel: ", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("from comma ok", v, ok)
			} else {
				fmt.Println("from quit with comma ok : ", v)
			}
			return
		}
	}
}

//send
func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// q <- false
	close(q)
	close(e)
	close(o)
}
