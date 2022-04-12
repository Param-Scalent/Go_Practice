package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("Sample.txt")
	if err != nil {
		log.Fatal("Error", err)
	}
	file.WriteString("Hi my is Param, and this file was created using GO!!")
	file.Close()

	stream, err := ioutil.ReadFile("Sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	s1 := string(stream)
	fmt.Println(s1)
}
