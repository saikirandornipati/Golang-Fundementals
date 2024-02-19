package main


import "fmt"


func main() { 

	// var ptr *string

	// fmt.Println(ptr)


	myNumber :=23
	var ptr =  &myNumber
	fmt.Println("value of actual pointer is ",*ptr)

	new_number := *ptr +23
	fmt.Println(new_number)
}