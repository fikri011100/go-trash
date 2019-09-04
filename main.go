package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	// using gorilla mux
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HelloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("Go running on port 8081")

	handleRequests()
}
