package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	//Comma Ok || Error Ok syntax (try/catch)
	input, _ := reader.ReadString('\n')
	// input, err := reader.ReadString('\n')
	// _, err := reader.ReadString('\n')

	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of this rating is %T", input)
}
