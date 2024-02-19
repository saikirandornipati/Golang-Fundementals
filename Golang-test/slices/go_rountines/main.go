package main

import (
	"fmt"
	"time"
)

func infinitelooping(thing string) {
	for i := 0; i < 10; i++ {
		fmt.Println(thing)
		time.Sleep(10)
	}

}

func main() {
	fmt.Println("hello world")
	go infinitelooping("dog")
	infinitelooping("cat")
}
