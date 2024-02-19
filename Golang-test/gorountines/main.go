package main

import ("fmt"
		"time")

func main(){
	// greeting("hello")
	// greeting("world")

}


// func greeting(s string){
// 	for i:=0;i<6;i++{
// 		time.Sleep(2 * time.Second)

// 		fmt.Println(s)
// 	}
// }


func getStatusCode(endpoint string){
	res, err := http.Get(endpoint)
	if err != nil{
		
	}
}