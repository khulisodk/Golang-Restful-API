package main

import (
	"encoding/json"   // Json format data
	"github.com/gorilla/mux"  // 3rd party router
	"log"
	"math/rand"  // Random number generation
	"net/http" // for http
	"strconv" // string converter
)


//Creating a book Struct/Model(can be a Class in C languages)
type Book struct {

	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Tittle string `json:"tittle"`
	Author *Author `json:"author"`

}

//creating an Author Struct/Model
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

//Initialise books as a slice of Book struct

var books []Book


//Creating Handlers Functions or end points
//Takes two parameters i.e. Response and Request


//Get All Books
func getBooks(writer http.ResponseWriter,request *http.Request)  {

	// Make sure that the content is in Json
	writer.Header().Set("Content-Type","application/json")
	json.NewEncoder(writer).Encode(books)

}

//Get Single Book
func getBook(writer http.ResponseWriter,request *http.Request)  {
	// Make sure that the content is in Json
	writer.Header().Set("Content-Type","application/json")

	 params :=mux.Vars(request) // Get Parameters

	 //Loop through books and find the requested id
	 // range to loop through our slice 'books'
	 for _,item := range books{
        //id in the url
	 	if item.ID ==params["id"]{
	 		  //passing our writer and output our item
			json.NewEncoder(writer).Encode(item)
			return
		}
	 }
	    json.NewEncoder(writer).Encode(&Book{})

}

//Create a New Book
func createBook(writer http.ResponseWriter,request *http.Request)  {
	// Make sure that the content is in Json
	writer.Header().Set("Content-Type","application/json")
    var book Book //Book Struct

	json.NewDecoder(request.Body).Decode(&book)

   //Create an id for a new book
   //Generate an int random number and convert it into string
   book.ID = strconv.Itoa(rand.Intn(10000000))

   //Add our new book into books
   books = append(books,book)

      //output a single book
	   json.NewEncoder(writer).Encode(book)

}

//Update Single Book
func updateBook(writer http.ResponseWriter,request *http.Request)  {

	// Make sure that the content is in Json format
	writer.Header().Set("Content-Type","application/json")

	//int  mux request variable
	params := mux.Vars(request)

	//range to loop through the books
	for index , item := range books{

		//Check the item by id
		if item.ID == params["id"]{

			books = append(books[:index],books[index+1:]...)

			var book Book //Book Struct

			json.NewDecoder(request.Body).Decode(&book)

			book.ID = params["id"]

			//Add our new book into books
			books = append(books, book)
			//output the update book
			json.NewEncoder(writer).Encode(book)
			return
		}

	}

	json.NewEncoder(writer).Encode(books)

}



//Delete a Single Book params["id"]
func deleteBook(writer http.ResponseWriter,request *http.Request)  {

	// Make sure that the content is in Json
	writer.Header().Set("Content-Type","application/json")

	//Loop through books and find the requested id
	params := mux.Vars(request)
	//range to loop through the books
	for index , item := range books{
		//Check the item by id
		if item.ID == params["id"]{

			//loops through slice of books to get delete a particular book
			books = append(books[:index],books[index+1:]...)
			break
		}

	}

	// Displays the rest of books after deleting
	json.NewEncoder(writer).Encode(books)

}


func main(){

	//initialise Router
	router :=mux.NewRouter()

	//Create Mock Data append = add

	books = append(books,Book{ID:"1",ISBN:"4422311",Tittle:"Goland Fundamentals",Author:&Author{Firstname:"Khuliso",Lastname:"Mudau"}})
	books = append(books,Book{ID:"2",ISBN:"4422312",Tittle:"Goland Advanced Dev Guide",Author:&Author{Firstname:"Khuliso",Lastname:"Mudau"}})
	books = append(books,Book{ID:"3",ISBN:"4422313",Tittle:"Goland Web Service Dev ",Author:&Author{Firstname:"Khuliso",Lastname:"Mudau"}})

	//Create a Route Handlers to establish-End Points/url path and the type of http method to use
	//this takes two parameters i.e. path and function name

	router.HandleFunc("/api/mybooks",getBooks).Methods("GET")
	router.HandleFunc("/api/mybooks/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/mybooks", createBook).Methods("POST")
	router.HandleFunc("/api/mybooks/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/mybooks/{id}", deleteBook).Methods("DELETE")

	//To run our server, takes address and Router name
	//inside our log tag for error handlings

	log.Fatal(http.ListenAndServe(":8001", router))
	//log.Fatal(http.ListenAndServe(":8070", router))

}

