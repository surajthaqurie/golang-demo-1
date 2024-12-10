package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"` //aliases
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json data")
	EncodeJson()

}

func EncodeJson() {

	lcoCourses := []course{
		{"RectJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"Web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hig123", nil},
	}

	//Package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
