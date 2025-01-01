package main

import (
	"encoding/json"
	"fmt"
)

// @Server sends
type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

// @Server receives
type SensorReading struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Information Info   `'json:"info"`
}

type Info struct {
	Description string `json:"des"`
}

func main() {
	fmt.Println("Server sends JSON")

	author := Author{Name: "Elliot Forbes", Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}

	fmt.Printf("%+v\n", book)

	byteArray, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

	fmt.Println("Server receives JSON")

	jsonString := `{"name":"battery sensor", "capacity":40, "time":"2025-01-01T19:07:28Z", "info":{
	"des":"a sensor reading"
	}}`

	// var reading map[string]interface{}s
	var reading SensorReading
	err = json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Printf("%+v\n", reading)

}
