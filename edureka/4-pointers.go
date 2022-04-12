package main

import "fmt"

func main() {
	x := 10

	fmt.Println(x)
	fmt.Println(&x)

	// changeValue(x)
	changeValue(&x)

	fmt.Println("Now x is: ", x)
}

func changeValue(x *int) {
	*x = 7
}
