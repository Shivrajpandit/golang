package main

import "fmt"

func variablesExample() {
	// Go is statically typed. You can declare a variable and its type.
	var anInteger int = 42
	fmt.Println("Integer:", anInteger)

	// You can also declare a variable without initializing it.
	// It will be assigned a "zero-value" (0 for int, "" for string, etc.).
	var aString string
	aString = "This is a string."
	fmt.Println("String:", aString)

	// The `:=` syntax is shorthand for declaring and initializing a variable.
	// Go will infer the type automatically.
	aFloat := 3.14159
	aBoolean := true
	fmt.Printf("Float: %f, Boolean: %t\n", aFloat, aBoolean)

	// Constants are declared with `const`.
	const pi = 3.14
	fmt.Println("Constant Pi:", pi)
}
