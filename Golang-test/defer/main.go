package main

import "fmt"


func main(){
	defer fmt.Println("world")
	defer fmt.Println("hi")
	defer fmt.Println("keka")
	fmt.Println("Hello")
	defering()	

}
func defering(){
	for i :=0;i<5;i++{
		defer fmt.Println(i)
	}
}