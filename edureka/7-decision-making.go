package main

import "fmt"

func main() {
	var age int = 18

	if age > 18 {
		fmt.Println("you can vote")
	} else {
		fmt.Println("No, You can't Vote!")
	}

	switch age {
	case 16:
		fmt.Println("Prepare for College")
	case 18:
		fmt.Println("Don't run after girls")
	case 20:
		fmt.Println("Get yourself a job")
	default:
		fmt.Println("Are you even alive?")
	}

}
