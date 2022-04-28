package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type filePrintA struct{}

func main() {
	args := os.Args
	var filePointer string

	fmt.Println(args)

	if args != nil && len(args) > 1 {
		filePointer = args[1]
	} else {
		fmt.Println("Please Provide the file path")
		os.Exit(1)
	}

	fileName, err := os.Open(filePointer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileName.Close()

	printFileA(fileName)
}

func getFilePath(args []string) (string, error) {
	if args != nil && len(args) > 1 {
		return args[1], nil
	}
	return "", errors.New("Please Provide the file path")
}

func printFileA(file *os.File) {
	io.Copy(os.Stdout, file)
	fmt.Println(file)
}
