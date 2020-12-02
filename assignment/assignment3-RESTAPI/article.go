package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	ID          int    `json:id`
	Title       string `json:title`
	Description string `json:description`
	Content     string `json:content`
}

var articles []Article

func main() {

	router := mux.NewRouter()
	articles = append(articles, Article{ID: 1, Title: "Title here", Description: "Describe your Article", Content: " Content of your Article"},
		Article{ID: 2, Title: "Title here", Description: "Describe your Article", Content: " Content of your Article"},
		Article{ID: 3, Title: "Title here", Description: "Describe your Article", Content: " Content of your Article"},
		Article{ID: 4, Title: "Title here", Description: "Describe your Article", Content: " Content of your Article"},
		Article{ID: 5, Title: "Title here", Description: "Describe your Article", Content: " Content of your Article"},
		Article{ID: 6, Title: "Title here", Description: "Describe your Article", Content: " Content of your Article"})

	//define routes
	router.HandleFunc("/articles", getArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/articles", addArticle).Methods("POST")
	router.HandleFunc("/articles", updateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", removeArticle).Methods("DELETE")

	//starting a server
	log.Fatal(http.ListenAndServe(":5050", router))
}
func getArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])

	for _, article := range articles {
		if article.ID == i {
			json.NewEncoder(w).Encode(&article)
		}
	}
}
func addArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	articles = append(articles, article)
	json.NewEncoder(w).Encode(articles)
}
func updateArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	json.NewDecoder(r.Body).Decode(&article)

	for i, item := range articles {
		if item.ID == article.ID {
			articles[i] = article
		}
	}
	json.NewEncoder(w).Encode(articles)
}
func removeArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, item := range articles {
		if item.ID == id {
			articles = append(articles[:i], articles[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(articles)
}
