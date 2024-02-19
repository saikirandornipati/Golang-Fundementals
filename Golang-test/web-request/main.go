package main

import (
	"fmt"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("go request")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	defer response.Body.Close()

}
