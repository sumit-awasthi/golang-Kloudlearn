package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "books_store"
)

var db *sql.DB
var err error

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	db = DB

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Book
	result, err := db.Query("SELECT id, title, author, year from books")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var book Book
		err := result.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO books(id, title, author, year) VALUES($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	id := keyVal["id"]
	title := keyVal["title"]
	author := keyVal["author"]
	year := keyVal["year"]
	_, err = stmt.Exec(id, title, author, year)
	if err != nil {
		panic(err.Error())
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, title, author, year FROM books WHERE id = $1", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var book Book
	for result.Next() {
		err := result.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newTitle := keyVal["title"]
	newAuthor := keyVal["author"]
	newYear := keyVal["year"]
	_, err = stmt.Exec(newTitle, newAuthor, newYear, params["id"])
	if err != nil {
		panic(err.Error())
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM books WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
}
