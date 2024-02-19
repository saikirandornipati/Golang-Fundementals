package main


import "fmt"


func main(){
	fmt.Println("loops in golang")


	days := []string{"sun","mon","tues","thursday"}

	fmt.Println(days)

	// for d:=0;d<len(days);d++{
	// 	fmt.Println(days[d])


	// }

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	for index, day:=range days{
		fmt.Println(index,day)
	}
}