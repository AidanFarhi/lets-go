package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	numPeople := 10_000_000
	// this runs much faster
	// personSlice := make([]Person, 0, numPeople)
	// this does not
	personSlice := make([]Person, 0, 0)
	for i := 0; i < numPeople; i++ {
		personSlice = append(personSlice, Person{"bob", "bobson", 20})
	}
}
