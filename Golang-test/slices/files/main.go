package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("welcome to the files")

	content := "Welcome to filesystem of golang.org"

	file, err := os.Create("/.welcome.text")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is :", length)
	defer file.Close()

}
