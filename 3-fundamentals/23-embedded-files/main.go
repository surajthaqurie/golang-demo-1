package main

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed test.txt
var s string

//go:embed assets/*
var assets embed.FS

func main() {
	fmt.Println("Embedded files/directory:")

	fmt.Println(s)

	fs := http.FileServer(http.FS(assets))

	http.ListenAndServe(":8000", fs)

}
