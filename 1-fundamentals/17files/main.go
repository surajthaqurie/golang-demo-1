package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")

	// Create a file
	file, err := os.Create("./myFile.txt")
	checkNilErr(err)

	// Write the content in the file
	content := "This is needs to go in a file -go.dev"
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is ", length)
	defer file.Close()

	readFile("./myFile.txt")

}

func readFile(filename string) {
	databytes, err := os.ReadFile(filename)

	checkNilErr(err)

	// fmt.Println("Text data inside the file is \n", databytes)
	fmt.Println("Text data inside the file is \n", string(databytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err) // panic: shutdown the program and show the error
	}
}
