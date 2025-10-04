package main

import "fmt"

func controlFlowExample() {
	// --- For Loop ---
	// Go has only one looping construct, the `for` loop.
	fmt.Println("Standard for loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// --- If-Else ---
	// Standard if-else statement.
	num := 10
	if num%2 == 0 {
		fmt.Printf("%d is an even number.\n", num)
	} else {
		fmt.Printf("%d is an odd number.\n", num)
	}

	// You can also have a short statement before the condition.
	// `v` is only in scope within the if-else block.
	if v := 3; v < 5 {
		fmt.Println("v is less than 5")
	}

	// --- Switch ---
	// A switch statement is a shorter way to write a sequence of if-else statements.
	day := "Sunday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	case "Monday":
		fmt.Println("Ugh, it's Monday.")
	default:
		fmt.Println("It's a weekday.")
	}
}
