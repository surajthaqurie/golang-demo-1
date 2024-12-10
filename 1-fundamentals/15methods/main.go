package main

import "fmt"

func main() {
	fmt.Println("Welcome to the struct in golang")

	//No inheritance in golan; no super or parent

	user := User{Name: "suraj", Email: "suraj.go.dev", Status: true, Age: 1}
	fmt.Println(user)
	fmt.Printf("User details are: %+v", user)
	fmt.Printf("User name is: %v and email is %v\n", user.Name, user.Email)

	user.GetStatus()
	user.NewMail()

	fmt.Printf("User name is: %v and email is %v\n", user.Name, user.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// COPY of User struct
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"

	fmt.Println("Email of this user is:", u.Email)
}
