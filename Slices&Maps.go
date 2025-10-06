package main

import "fmt"

func slicesAndMapsExample() {
	// --- Slices ---
	// Slices are a key data type in Go, giving a more powerful
	// interface to sequences than arrays.
	s := make([]string, 3)
	fmt.Println("Empty slice:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set:", s)
	fmt.Println("Get s[2]:", s[2])
	fmt.Println("Length:", len(s))

	// Append items to a slice.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Appended:", s)

	// Slices can be sliced!
	l := s[2:5]
	fmt.Println("Slice s[2:5]:", l)

	// --- Maps ---
	// Maps are Go's built-in associative data type (sometimes called
	// hashes or dicts in other languages).
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("Map:", m)

	v1 := m["k1"]
	fmt.Println("Value for k1:", v1)

	// Delete a key-value pair from a map.
	delete(m, "k2")
	fmt.Println("Map after deleting k2:", m)

	// When getting a value from a map, you can also check if the key was present.
	_, prs := m["k2"]
	fmt.Println("Is k2 present?:", prs)
}
