package main

import "fmt"

func conditionals() {
	age := 11       // integer variable
	day := "Sunday" // string variable

	// IF statement example
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	// SWITCH statement example
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Weekend is near")
	case "Sunday":
		fmt.Println("It's a holiday")
	default:
		fmt.Println("Just a normal day")
	}
}
