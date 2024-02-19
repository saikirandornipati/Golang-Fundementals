package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://loc.dev"

func main() {

	fmt.Println("Starting")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is: %T\n", response)

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(data)
	fmt.Println(content)

	fmt.Println("Closing")

}
