package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db := InitDb()
	defer db.Close()

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Create User Called")

	db := InitDb()
	defer db.Close()

	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)

	json.NewEncoder(w).Encode(&user)
	//json.NewEncoder(w).Encode("Success create user")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Called")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Called")
}
