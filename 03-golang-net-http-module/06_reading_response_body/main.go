package main

func main () {
	url := "https://jsonplaceholder.typicode.com/todos"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close() // Prevents any kind of resource leaks

}