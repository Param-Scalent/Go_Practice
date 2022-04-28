package main

import (
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}
type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more tea",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran - ", e)
}
