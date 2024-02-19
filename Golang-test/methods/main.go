package main


import "fmt"


func main(){

	kiran := User{"kiran","saikiran@google.com",true,155,"Btech"}

	fmt.Println(kiran)
	fmt.Println("name is %v",kiran.Name)
	kiran.GetStatus()
	kiran.Newmail()

}


type User struct {
	Name string
	Email string
	Status bool
	Age int
	Degree string

}


func (u User) GetStatus(){
	fmt.Println("is user is avtive:",u.Status)


}

func (u User) Newmail(){
	u.Email = "saikiran@microsoft.com"
	fmt.Println(u.Email)

}


