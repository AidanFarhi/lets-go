package main

import "fmt"

// using embedding for composition
type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee // Employee is embedded
	Reports  []Employee
}

func (m Manager) HasReport(e Employee) bool {
	for _, r := range m.Reports {
		if r.ID == e.ID {
			return true
		}
	}
	return false
}
