package main

import (
	"errors"
	"fmt"
)

// In Go, it's idiomatic to communicate errors via an explicit,
// separate return value. This contrasts with exceptions in other languages.
// By convention, the error is the last return value.
func divide(a, b float64) (float64, error) {
	// If the divisor is zero, we return an error.
	if b == 0 {
		// `errors.New` constructs a basic error value with the given message.
		return 0, errors.New("cannot divide by zero")
	}
	// If no error, we return the result and a `nil` error value.
	return a / b, nil
}

func errorHandlingExample() {
	// Test case 1: Successful division
	result, err := divide(10.0, 2.0)
	if err != nil {
		// This block won't execute because there is no error.
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("10.0 / 2.0 =", result)
	}

	// Test case 2: Division by zero
	result, err = divide(10.0, 0.0)
	if err != nil {
		// This block WILL execute.
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("10.0 / 0.0 =", result)
	}
}
