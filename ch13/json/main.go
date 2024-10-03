package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {

	// read the data from the files as bytes
	data, err := os.ReadFile("orders.json")
	if err != nil {
		handleError(err)
	}

	// unmarshal the data into a slice
	var orders []Order
	err = json.Unmarshal(data, &orders)
	if err != nil {
		handleError(err)
	}

	// extract all the items and write them to a file
	output, err := os.Create("items.json")
	defer output.Close()
	if err != nil {
		handleError(err)
	}
	var items []Item
	for _, o := range orders {
		for _, i := range o.Items {
			items = append(items, i)
		}
	}
	encoder := json.NewEncoder(output)
	encoder.Encode(items)
}
