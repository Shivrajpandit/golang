package main

import "fmt"

// A simple function that takes two integers and returns their sum as an integer.
func add(a int, b int) int {
	return a + b
}

// Go functions can return multiple values. This is often used to return
// a result and an error value.
func swap(x, y string) (string, string) {
	return y, x
}

func functionsExample() {
	// Calling the `add` function.
	result := add(10, 20)
	fmt.Println("10 + 20 =", result)

	// Calling a function that returns multiple values.
	a, b := "hello", "world"
	fmt.Println("Before swap:", a, b)
	a, b = swap(a, b)
	fmt.Println("After swap:", a, b)
}
