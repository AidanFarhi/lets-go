package main

import (
	"fmt"
	"time"
)

// a user defined type with an underlying type of struct
type Counter struct {
	total       int
	lastUpdated time.Time
}

// method with a pointer receiver
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// method with a value receiver
func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated %v", c.total, c.lastUpdated)
}
