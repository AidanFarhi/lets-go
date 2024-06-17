package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	data := Response{}
	url := "https://jsonplaceholder.typicode.com/todos/1"
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		url,
		nil,
	)
	req.Header.Add("X-My-Client", "Learning Go")
	res, _ := client.Do(req)
	defer res.Body.Close()
	fmt.Println(res.Header.Get("Content-Type"))
	json.NewDecoder(res.Body).Decode(&data)
	fmt.Printf("%v\n", data)
}
