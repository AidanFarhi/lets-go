package main

import "fmt"

type Logger interface {
	Log(message string)
}

// A function type with a method on it
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// A simple utility function for logging
func LogOutput(message string) {
	fmt.Println(message)
}
