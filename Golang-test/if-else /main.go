package main

import "fmt"



func main(){
	loginCount :=10
	var result string


	if loginCount > 10{
		result = "regular"
	}else if loginCount<10{
		result = "non user"
	}else{
		result = "this is okay"
	}
	fmt.Println(result)


	
}