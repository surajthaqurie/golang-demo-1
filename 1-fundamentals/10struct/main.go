package main

import "fmt"

func main() {
	fmt.Println("Welcome to the struct in golang")

	//No inheritance in golan; no super or parent

	user := User{Name: "suraj", Email: "suraj.go.dev", Status: true, Age: 1}
	fmt.Println(user)
	fmt.Printf("User details are: %+v", user)
	fmt.Printf("User name is: %v and email is %v", user.Name, user.Email)

}

// If first letter is Capital than that is exportable(public) other not (private)
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
