package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!!!!")
	})

	fmt.Println("Debugging")
	fmt.Println("0")
	fmt.Println("1")

	http.ListenAndServe(":8080", r)
}
