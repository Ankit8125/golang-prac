package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CatFactResponse struct {
	Fact string `json:"fact"`
	Length int `json:"length"`
}

func main () {
	url := "https://catfact.ninja/fact"
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

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body failed", err)
		return
	}

	// JSON unmarshal is going to convert bytes into a struct (our response format). (or you have to create a model of the data that you are going to receive)
	var data CatFactResponse

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		fmt.Println("json unmarshal failed", err)
		return
	}

	fmt.Println(data.Fact, data.Length)

}