package main

import (
	"fmt"
	"net/http"
)

func main () {
	url := "https://jsonplaceholder.typicode.com/todos"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close() // Prevents any kind of resource leaks	

	fmt.Println("status code ", resp.StatusCode)
	fmt.Println("status ", resp.Status)
}