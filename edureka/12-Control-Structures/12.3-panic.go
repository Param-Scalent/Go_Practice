package main

import "fmt"

func main() {

	fmt.Println(addemup(10, 20, 30, 40, 50))

}

func addemup(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}
