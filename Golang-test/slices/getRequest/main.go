package main

import (
	"encoding/json"
	"fmt"

	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Books struct {
	Id     string  "json"
	Name   string  `json:"name"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:website`
}

func main() {
	fmt.Println("Web get request")

}

var books []Books

// middleware
func (b *Books) IsEmpty() bool {
	return b.Name == ""

}

//controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>check the api,its working</h1>"))
}

func getAllbooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all books")
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func getOnebook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one book")
	w.Header().Set("content-Type", "application/json")

	params := mux.vars(r)
	for _,book := range books{
		if books.Id == params["id"]{
			json.NewEncoder(w).Encode(book)
		}
	}
	json.NewEncoder(w).Encode("no book is not available")
	return 

}


func createOnebook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a book")
	w.Header().Set("content-Type", "application/json")
	if r.Body == nil{
		json.NewEncoder(w).Encode("please send some data")

	}
	var book Book 
	_= json.NewDecoder(r.Body).Decode(&book)
	if book.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside json")
		return 
	}

	//generating the book id at backend

	rand.Seed(time.Now().unixNano())
	book.Id = strconv.Itoa(rand.Intn(100))
	json.NewEncoder(w).Encode(book)
	return 





}
