package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("Welcome to switch case")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Print("value of dice is", diceNumber)

	switch diceNumber {

	case 1:
		fmt.Print("Dice value in 1 you can open")
		fallthrough // Run another case also

	case 2:
		fmt.Print("You can move 2 spot")
		fallthrough

	case 3:
		fmt.Print("You can move 3 spot")
		fallthrough

	case 4:
		fmt.Print("You can move 4 spot")
		fallthrough

	case 5:
		fmt.Print("You can move 5 spot")

	case 6:
		fmt.Print("You can move to 6 stop and roll dice again")

	default:
		fmt.Print("What was that")

	}

}
