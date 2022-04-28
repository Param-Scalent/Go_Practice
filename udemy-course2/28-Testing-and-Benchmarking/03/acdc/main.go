package main

import (
	"fmt"
	"udemy/course2/28-Testing-and-Benchmarking/03/acdc"
)

func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}
