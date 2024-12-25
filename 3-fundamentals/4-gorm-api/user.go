package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")

	}

	defer db.Close()

	db.AutoMigrate(&User{})

}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all user endpoint hit")

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database")

	}

	defer db.Close()

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New user endpoint hit")

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database")

	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	fmt.Println(w, "New user successfully created")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete user endpoint hit")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database")

	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name= ?", name).Find(&user)
	db.Delete(&user)

	fmt.Println(w, "User successfully deleted")

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update user endpoint hit")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database")

	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name= ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Println(w, "User successfully updated")

}
