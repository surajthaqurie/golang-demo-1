package main

import "fmt"

func main() {
	fmt.Println("Welcome to the maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	fmt.Println("Js shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages", languages)

	//Loop are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v,value is %v\n", key, value)
	}

}
