package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args[1])

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
