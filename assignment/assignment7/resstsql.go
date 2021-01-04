package main

// REst api using sql connection
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

// Product struct
type Product struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

// db is an object of sql connection
var db *sql.DB
var err error

func main() {
	// DB Setup
	db, err = sql.Open("mysql", "root:root@(127.0.0.1:3306)/mysqlcurd?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected")
	defer db.Close()

	// Init Router, here we are using gorilla mux router
	router := mux.NewRouter()

	// Route Handler which establish endpoints

	router.HandleFunc("/api/products", createProduct).Methods("POST")
	router.HandleFunc("/api/products", getProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	router.HandleFunc("/api/products/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", deleteProduct).Methods("DELETE")

	// Run Server
	log.Fatal(http.ListenAndServe(":8000", router))
}

// createProduct create a new product
func createProduct(w http.ResponseWriter, r *http.Request) {

	statement, err := db.Prepare("INSERT INTO product(id,isbn,name,price)VALUES(?,?,?,?)")
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
	name := dataMap["name"]
	price := dataMap["price"]

	_, err = statement.Exec(id, isbn, name, price)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New product added")
}

// getProduct return a single product based on id which you pass.
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("SELECT * FROM product WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	// product is object of Book struct
	var product Product

	for result.Next() {
		err := result.Scan(&product.ID, &product.Isbn, &product.Name, &product.Price)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(product)
}

// getProducts returns all products.
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product

	result, err := db.Query("SELECT * FROM product")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var product Product
		err := result.Scan(&product.ID, &product.Isbn, &product.Name, &product.Price)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)

	}
	json.NewEncoder(w).Encode(products)
}

// updateProduct update the product name based on id passed.
func updateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	statement, err := db.Prepare("UPDATE products SET name = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	dataMap := make(map[string]string)
	json.Unmarshal(body, &dataMap)
	newTitle := dataMap["name"]

	_, err = statement.Exec(newTitle, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Product %s updated", params["id"])

}

// deleteProduct delete a product based on id
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	statement, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = statement.Exec(params["id"])
	if err != nil {
		panic(err.Error)
	}
	fmt.Fprintf(w, "Product %s deleted", params["id"])
}
