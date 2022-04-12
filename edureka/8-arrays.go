package main

import "fmt"

func main() {
	var EvenNum [5]int

	EvenNum[0] = 2
	EvenNum[1] = 4
	EvenNum[2] = 6
	EvenNum[3] = 8
	EvenNum[4] = 10

	fmt.Println(EvenNum[2])

	OddNum := [5]int{1, 3, 5, 7, 9}

	fmt.Println(OddNum[2])

	for i, v := range OddNum {
		fmt.Println(i, v)
	}

	// Slicing arrays
	numSlice := []int{5, 4, 3, 2, 1}

	sliced := numSlice[3:5]

	fmt.Println("Sliced arrays: ", sliced)

	slice2 := make([]int, 5, 10)

	copy(slice2, numSlice)

	fmt.Println("Copied slices: ", slice2)
	fmt.Println("Copied from: ", numSlice)

	slice3 := append(numSlice, 3, 0, 1)
	fmt.Println("Appended Slice: ", slice3)
}
