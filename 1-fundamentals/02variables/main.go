package main

import "fmt"

// jwtToken := 30000 // not worked
// var jwtToken int = 30000 (worked)

const LoginToken string = "tupxwko"

func main() {
	// String
	var username string = "suraj"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n ", username)

	//Boolean
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n ", isLoggedIn)

	//Numeric types
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n ", smallVal)

	//Float types
	var smallFloat float32 = 255.42354654864789
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n ", smallFloat)

	//Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n ", anotherVariable)

	//Implicit types
	var website = "nepal.com"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n ", website)

	//No var style
	numberOfUser := 30000.2
	fmt.Println(numberOfUser)
	fmt.Printf("variable is of type: %T \n ", numberOfUser)

	//Public
	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n ", LoginToken)

}
