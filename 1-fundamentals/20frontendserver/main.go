package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to the server for frontend in golang")

	PerformGetRequest()

	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code", response.Status)
	fmt.Println("Content length is", response.ContentLength)

	// content, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)

	// }
	// fmt.Println("Content", string(content))

	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)

	}

	byteCount, err := responseString.Write(content)
	if err != nil {
		panic(err)

	}
	fmt.Println("Byte count is:", byteCount)
	fmt.Println("Response string", responseString.String())

}

func PerformPostJsonRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts"

	requestBody := strings.NewReader(`
	{
		"title": "foo",
		"body": "bar",
		"userId": 1
	}
`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)

	}

	fmt.Println("Content", string(content))

}
