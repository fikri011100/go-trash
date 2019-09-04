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

	// users endpoint
	router.HandleFunc("/users", AllUser).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", UpdateUser).Methods("PATCH")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("Go running on port 8081")

	handleRequests()
}
