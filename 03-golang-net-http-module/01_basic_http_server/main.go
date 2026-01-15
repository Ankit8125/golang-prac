package main

import (
	"fmt"
	"net/http"
)

func helloHandler (w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet { // Checks if the request method is NOT GET:
		http.Error(w, "only GET is allowed", http.StatusMethodNotAllowed) // hover to see what all params it takes
		return 
	}

	_, _ = w.Write([] byte ("Hello from GO net/http server")) // Writes the response body to the client. Converts the string to a byte slice [] byte (...)
	// _, _ ignores the return values (number of bytes written and any error)
}

func main () {
	// Registering a route
	http.HandleFunc("/hello", helloHandler)
	// func http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	// - ResponseWriter is going to write the response back to the client. (GPT: An interface used to write the HTTP response back to the client)
	// - Request - this is the incoming request info that is going to have our headers, body, method, etc (GPT: A pointer to the request object containing method, headers, body, URL, etc.)

	fmt.Println("Going to port 8080")

	// Listening a particular port
	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
}