package main

import "fmt"

func main() {
	const pi float64 = 3.14642732
	var name string = "Aryya PAul"
	win := true
	x := 5

	fmt.Println(len(name))
	fmt.Println(name + " is a chill dude")

	fmt.Printf("%.3f \n", pi) // for precision to 3
	fmt.Printf("%T \n", name) // for Type
	fmt.Printf("%t \n", win)  // for booleans
	fmt.Printf("%d \n", x)    // fot int
	fmt.Printf("%b \n", 8)    // for printing binary
	fmt.Printf("%c \n", 34)   // for character asscoiated with key codes
	fmt.Printf("%x \n", 15)   // for hex codes
	fmt.Printf("%e \n", pi)   // for scientific notations
}
