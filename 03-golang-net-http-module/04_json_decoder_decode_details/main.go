package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Encoding the data to JSON and writing to ResponseWriter
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data) // this is going to send the data back 
}

type TestRequest struct {
	Name string `json:"name"` // `json` map the JSON field names to the go struct field. (No space after ":")
}
// TestRequest = the struct name 
// Name = a field that holds a string `json:"name"` = a tag that maps JSON field to Go field
// When JSON has {"name": "ankit"}, the decoder puts "ankit" in the Name field. Without this tag, the decoder wouldn't know which JSON field maps to which Go field

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // Checks if the request is a POST request.
		writeJSON(w, http.StatusMethodNotAllowed, map[string] any {
			"ok": false,
			"error": "only POST is allowed",
		})
		return
	}
	defer r.Body.Close() // This is best practice to avoid memory leaks.
	// r.Body contains the raw JSON data the client sent
	// .Close() releases the resource (cleanup) after we're done reading it

	var req TestRequest // Creating a Variable to Hold Decoded Data:

	dec := json.NewDecoder(r.Body) // Creates a JSON decoder that reads from r.Body (the request body).

	if err := dec.Decode(&req); err!= nil {
		// Decodes JSON from request body into the req struct. Client sends: {"name": "ankit"}. Decoder reads it and fills: req.Name = "ankit"
		// &req = pointer to req (Decode needs a pointer to modify the struct)
		writeJSON(w, http.StatusBadRequest, map [string] any {
			"ok": false,
			"error": "Invalid JSON format",
		})
		return
	}

	/*
	dec.Decode(&req) does the following:
	- Reads the JSON string from r.Body (Raw text: {"name": "ankit"})
	- Parses it (breaks it into parts) (Key: "name", Value: "ankit")
	- Looks at the struct definition: (type TestRequest struct { Name string `json:"name"`} )
	- The decoder sees: json:"name" tag
	- This tag says: "When you see JSON key 'name', put its value in the Go field 'Name'"
	- Matches and fills:
	*/

	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map [string] any {
			"ok": false,
			"error": "name must not be empty!",
		})
		return 
	}

	writeJSON(w, http.StatusOK, map[string] any {
		"ok": true,
		"data": req,
		"timeStamp": time.Now().UTC(),
	})
}

// Decoder is going to read the JSON from my request to body
func main () {
	// Receiving JSON request using json.Decoder
	http.HandleFunc("/test", testHandler)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println("error: ", err)
}