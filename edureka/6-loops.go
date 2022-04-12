package main

import "fmt"

func main() {
	fmt.Println("-----------For Loop--------------")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("-----------While Loop--------------")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 1; i <= 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
