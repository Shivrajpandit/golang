package main

import (
	"fmt"
	"net/http"
)

// A handler function for incoming web requests.
// It must have this exact signature.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// `r` contains information about the request (URL, headers, etc.).
	// `w` is where you write your response to.
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func SimpleWebServer() {
	// Register the handler function for the "/" route.
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server starting on port 8080...")
	fmt.Println("Open http://localhost:8080 in your browser.")

	// Start the web server on port 8080.
	// `ListenAndServe` blocks until the server is stopped.
	// If it returns an error, we log it.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
