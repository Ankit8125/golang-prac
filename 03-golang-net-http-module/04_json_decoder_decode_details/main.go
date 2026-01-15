package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data) // this is going to send the data back 
}

type TestRequest struct {
	Name string `json: "name"` // `json` map the JSON field names to the go struct field
}

func testHandler(w http.ResponseWriter, r* http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string] any {
			"ok": false,
			"error": "only POST is allowed",
		})

		return
	}
	defer r.Body.Close()

	var req TestRequest

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&req); err!= nil {
		writeJSON(w, http.StatusBadRequest, map [string] any {
			"ok": false,
			"error": "Invalid JSON format"
		})

		return
	}

	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map [string] any {
			"ok": false,
			"error": "name must not be empty!"
		})
		return 
	}

	writeJSON(w, http.StatusOK, map[string] any {
		"ok": true,
		"data": req,
		"timeStamp": time.Now()
	})
}

// Decoder is going to read the JSON from my request to body
func main () {

	http.HandleFunc("/test", testHandler)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println("error: ", err)
}