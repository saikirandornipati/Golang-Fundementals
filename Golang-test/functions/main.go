package main


import "fmt"

func main(){
	fmt.Println("welcome to functions")
	fmt.Println(checkingVariable(10,20))
	proResult := proAdder(2,3,4)
	fmt.Println(proResult)

}

func checkingVariable(a int, b int ) int {
	return a+b
}

func proAdder(values ...int)int{
	total:=0
	for _,val :=range values{
		total += val
	}

	return total
}