package main

import (
	"fmt"
	"io"
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

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body) // "io" is going to read all our body bytes that we are going to receive

	if err != nil {
		fmt.Println(err)
		return
	}

	bodyText := string(bodyBytes)
	max := 250

	if len(bodyText) < max {
		max = len(bodyText)
	}

	fmt.Println(bodyText[:max])
}