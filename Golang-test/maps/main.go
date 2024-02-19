package main


import "fmt"


func main(){
	fmt.Println("maps in golang")

	languages := make(map[int]string)

	languages[1] = "js"
	languages[2] = "ruby"
	fmt.Println(languages[1])


	//languages := make(map[int]string)
	languages[1] = "aaa"
	languages[2] = "ruby"
	fmt.Println(languages[1])


	for _,value :=range languages{
		fmt.Println(value)
	}
}