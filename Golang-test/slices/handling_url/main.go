package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev"

func main() {

	fmt.Println("welcome to the urls package")

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())

}
