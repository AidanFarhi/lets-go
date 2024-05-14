package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

type Sentinal string

func (s Sentinal) Error() string {
	return string(s)
}

type EmptyFieldErr struct {
	Field   string
	Message string
}

func (efe EmptyFieldErr) Error() string {
	return efe.Message
}

const (
	ErrInvalidID = Sentinal("invalid id")
)

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

func ValidateEmployee(e Employee) error {
	if len(e.ID) == 0 {
		return EmptyFieldErr{
			Field:   "ID",
			Message: "empty ID field",
		}
	}
	if !validID.MatchString(e.ID) {
		return ErrInvalidID
	}
	if len(e.FirstName) == 0 {
		return EmptyFieldErr{
			Field:   "FirstName",
			Message: "empty FirstName field",
		}
	}
	if len(e.LastName) == 0 {
		return EmptyFieldErr{
			Field:   "LastName",
			Message: "empty LastName field",
		}
	}
	if len(e.Title) == 0 {
		return EmptyFieldErr{
			Field:   "Title",
			Message: "empty Title field",
		}
	}
	return nil
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	errorsFound := make([]string, 0, 0)
	validRecordsFound := make([]string, 0, 0)
	for d.More() {
		var emp Employee
		var emptyFieldErr EmptyFieldErr
		err := d.Decode(&emp)
		err = ValidateEmployee(emp)
		if errors.Is(err, ErrInvalidID) {
			errorsFound = append(errorsFound, "invalid id")
			continue
		}
		if errors.As(err, &emptyFieldErr) {
			errorsFound = append(errorsFound, "empty field provided: "+emptyFieldErr.Field)
			continue
		}
		count++
		validRecordsFound = append(validRecordsFound, fmt.Sprintf("record %d: %+v", count, emp))
	}
	fmt.Println("--- Valid Records ---")
	for _, rec := range validRecordsFound {
		fmt.Println(rec)
	}
	fmt.Println("--- Errors ---")
	for _, e := range errorsFound {
		fmt.Println(e)
	}
}
