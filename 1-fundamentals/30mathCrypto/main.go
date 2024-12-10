package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to the math crypto and random number in the golang")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4
	// fmt.Println("The sum is:", myNumberOne+myNumberTwo)

	// Random number Math
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// Random number Crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
