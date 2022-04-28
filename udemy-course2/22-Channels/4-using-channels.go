package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive
	bar(c)
	fmt.Println("about to exit")
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
