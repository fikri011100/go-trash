package handler

import (
	"encoding/json"
	"fmt"
	"github.com/zufaralam02/first-go/db"
	"github.com/zufaralam02/first-go/model"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	defer db.Close()

	var users []model.User
	db.Find(&users)

	//json.NewEncoder(w).Encode(users)

	//users := []model.User{
	//	{Name: "Allam", Email: "zufaralam02@gmail.com"},
	//	{Name: "Zufar", Email: "allam.it@langitpayment.com"},
	//}

	responseJson(w, http.StatusOK, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Create User Called")

	db := db.InitDb()
	defer db.Close()

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)

	//json.NewEncoder(w).Encode(&user)
	//json.NewEncoder(w).Encode("Success create user")
	responseJson(w, http.StatusOK, "success create")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Called")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Called")
}
