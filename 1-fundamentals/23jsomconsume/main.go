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
	fmt.Println("Welcome to the json consume in golang")

	DecodeJson()

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
        {
            "name": "RectJS Bootcamp",
            "Price": 299,
            "website": "LearnCodeOnline.in",
            "tags": [ "Web-dev", "js" ]
        }
       
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)

		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//Some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	// fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
