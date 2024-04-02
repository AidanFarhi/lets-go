package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	persons := make([]Person, 0, 100_000_000) // this runs much faster
	// persons := []Person{}
	n := 100_000_000
	for i := 0; i < n; i++ {
		persons = append(persons, Person{"Bob", "Smith", 34})
	}
}
