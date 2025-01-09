package main

import (
	"fmt"
	"portscanner/port"
)

func main() {

	fmt.Println("Port scanner")
	open := port.ScanPort("tcp", "localhost", 8000)
	fmt.Printf("Port %s is %s\n", open.Port, open.State)

	results := port.InitialScan("localhost")
	fmt.Println(results)

}
