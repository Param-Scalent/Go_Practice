package main

import "fmt"

func main() {

	x, y := 5, 6

	fmt.Println("x + y: ", x+y)
	fmt.Println("x - y: ", x-y)
	fmt.Println("x * y: ", x*y)
	fmt.Println("x / y: ", x/y)
	fmt.Println("x % y: ", x%y)

	isBool := true

	hate := false

	fmt.Println(isBool && hate)
	fmt.Println(isBool || hate)
	fmt.Println(!isBool)
}
