package main

import (
	"fmt"
	"net/http"
)

func AllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users Called")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create User Called")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Called")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Called")
}
