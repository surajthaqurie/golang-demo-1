package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Welcome to handle url")
	fmt.Println(myurl)

	// Parsing the url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of params is %T\n", qparams)

	fmt.Println(qparams["name"])
	for _, val := range qparams {
		fmt.Println("Param is", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "jsonplaceholder.typicode.com",
		Path:    "posts",
		RawPath: "name=english",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another url", anotherUrl)
}
