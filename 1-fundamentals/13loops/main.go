package main

import "fmt"

func main() {
	fmt.Println("Welcome to the loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("Index is and value is %v\n", day)
	// }

	rogueValue := 1
	for rogueValue < 10 {

		if rogueValue == 2 {

			goto lco
		}

		if rogueValue == 5 {
			rogueValue++

			// break  // entire loop breakout
			continue
		}

		fmt.Println("Value", rogueValue)

		rogueValue++
	}

lco:
	fmt.Println("Jumping at golan.dev")

}
