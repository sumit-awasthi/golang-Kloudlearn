package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to new server!")
	})

	http.HandleFunc("/Trainee", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome Trainees!")
	})

	http.HandleFunc("/Employee", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome Employees!")
	})

	http.ListenAndServe(":8080", nil)
}
