package main

import (
	"io"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome!")
}

func howru(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hey whatsup!")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/hey whatsup", howru)

	http.ListenAndServe(":5050", mux)
}
