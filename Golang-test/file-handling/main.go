package main


import ("fmt"
		"io"
		"ioutil"
		"os"



)
func main(){
	fmt.Println("hello")

	content := "time is good"

	file, err  := os.Create("./myGolanfile.txt")

	if err != nil{
		panic(err)
	}
	length, err :=io.WriteString(file,content)

	if err !=nil{
		panic(err)


	}
	readFile("./myGolanfile.txt")

	fmt.Println(length)
	defer file.Close()
}




func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	// if err != nil{
	// 	panic(err)
	// }
	checkErr(err)

}


def checkErr(err error){
	if err != nil{
		panic(err)
	}

}