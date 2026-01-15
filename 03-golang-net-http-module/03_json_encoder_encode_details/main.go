package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler (w http.ResponseWriter, r *http.Request) {
	// Encode is writing the JSON and returning that
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := map [string] any {
		"ok": true,
		"message": "JSON encoded successfully",
		"datetime": time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)
}

func main () {

	http.HandleFunc("/ok", successHandler)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println("error: ", err)

}