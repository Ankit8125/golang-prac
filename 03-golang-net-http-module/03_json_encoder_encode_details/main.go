package main

import (
	"encoding/json" // Package used for encoding JSON data
	"fmt"
	"net/http"
	"time"
)

func successHandler (w http.ResponseWriter, r *http.Request) {
	// Encode is writing the JSON and returning that

	// How to set headers ?
	w.Header().Set("Content-Type", "application/json") // HTTP headers tell the client what type of data you're sending.
	// Content-Type: application/json = "Hey client, I'm sending you JSON data, not HTML or plain text". Without this, the browser might not recognize it as JSON
	
	// How to set a status code ?
	w.WriteHeader(http.StatusOK) // Sets the HTTP status code (if we do not write, then by default it is 200 OK)

	res := map [string] any { // This is Go's way of creating JSON-like data structures
		"ok": true,
		"message": "JSON encoded successfully",
		"datetime": time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res) // Converting Go data to JSON
}

func main () {
	// Sending JSON response using json.Encoder

	http.HandleFunc("/ok", successHandler)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println("error: ", err)

}