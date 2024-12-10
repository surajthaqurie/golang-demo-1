package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcome to form data in golang")

	PerformPostFormExample()
}

func PerformPostFormExample() {

	const myUrl = "https://jsonplaceholder.typicode.com/posts"

	//@Form Data (image)
	data := url.Values{}

	data.Add("Firstname", "suraj")
	data.Add("lastname", "chand")
	data.Add("email", "suraj@go.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
