package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handle route using handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hey Pal Welcome to new server! by sumit")
	})

	// listen to port
	http.ListenAndServe(":8080", nil)
}
