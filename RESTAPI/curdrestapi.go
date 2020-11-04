package main

// This is Simplest REST API Book CRUD example in Golang
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// db is an object of sql connection
var db *sql.DB
var err error

func main() {
	// DB Setup
	db, err = sql.Open("mysql", "root:password123@(127.0.0.1:3306)/mysqlcrud?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Init Router, here we are using gorilla mux router
	router := mux.NewRouter()

	// Route Handler which establish endpoints

	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run Server
	log.Fatal(http.ListenAndServe(":8080", router))
}

// createBook create a new book
func createBook(w http.ResponseWriter, r *http.Request) {

	statement, err := db.Prepare("INSERT INTO books(id,isbn,title,author)VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	dataMap := make(map[string]string)
	json.Unmarshal(body, &dataMap)
	id := dataMap["id"]
	isbn := dataMap["isbn"]
	title := dataMap["title"]
	author := dataMap["author"]

	_, err = statement.Exec(id, isbn, title, author)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New book created")
}

// getBook return a single book based on id which you pass.
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("SELECT * FROM books WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	// book is object of Book struct
	var book Book

	for result.Next() {
		err := result.Scan(&book.ID, &book.Isbn, &book.Title, &book.Author)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(book)
}

// getBooks returns all books.
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Book

	result, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var book Book
		err := result.Scan(&book.ID, &book.Isbn, &book.Title, &book.Author)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)

	}
	json.NewEncoder(w).Encode(books)
}

// updateBook update the Book title based on id passed.
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	statement, err := db.Prepare("UPDATE books SET title = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	dataMap := make(map[string]string)
	json.Unmarshal(body, &dataMap)
	newTitle := dataMap["title"]

	_, err = statement.Exec(newTitle, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Book %s updated", params["id"])

}

// deleteBook delete a book based on id
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	statement, err := db.Prepare("DELETE FROM books WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = statement.Exec(params["id"])
	if err != nil {
		panic(err.Error)
	}
	fmt.Fprintf(w, "Book %s deleted", params["id"])
}
