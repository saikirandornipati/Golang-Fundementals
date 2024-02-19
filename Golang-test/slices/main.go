package main	


import "fmt"


func main(){
	var fruitlist = []string{"apple","tamato","green"}
	fmt.Println("fruitlist is %T\n",fruitlist)


	fruitlist = append(fruitlist,"map","banana")
	fmt.Println(fruitlist[1])

	var courses = []string{"a","b","c","d","e"}

	fmt.Println(courses)
	var index int = 2

	courses = append(courses,[:index])

}