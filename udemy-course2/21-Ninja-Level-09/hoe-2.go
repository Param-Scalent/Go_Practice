package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak() string
}

func (p *person) speak() string {
	return "Hello from speak()"
}

func saySomething(h human) {
	fmt.Println("Call to speak", h.speak())
}

func main() {
	p := person{first: "James"}
	saySomething(&p)
}
