package main

import "net/http"

func main() {
	logger := LoggerAdapter(LogOutput)
	dataStore := NewSimpleDataStore()
	logic := NewSimpleLogic(logger, dataStore)
	controller := NewController(logger, logic)
	http.HandleFunc("/hello", controller.SayHello)
	http.ListenAndServe(":8080", nil)
}
