package main


import "fmt"


func main(){

	kiran := User{"kiran","saikiran@google.com",true,155}

	fmt.Println(kiran)
	fmt.Println("name is %v",kiran.Name)

}


type User struct {
	Name string
	Email string
	Status bool
	Age int 
}