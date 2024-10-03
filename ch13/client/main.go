package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Data struct {
	UserID    int64  `json:"userId"`
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"https://jsonplaceholder.typicode.com/todos/1",
		nil,
	)
	if err != nil {
		handleError(err)
	}
	req.Header.Add("X-My-Client", "My Client")
	res, err := client.Do(req)
	if err != nil {
		handleError(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		handleError(fmt.Errorf("unexpected status code: %d", res.StatusCode))
	}
	fmt.Println(res.Header.Get("Content-Type"))
	decoder := json.NewDecoder(res.Body)
	var data Data
	decoder.Decode(&data)
	fmt.Println(data)
}
