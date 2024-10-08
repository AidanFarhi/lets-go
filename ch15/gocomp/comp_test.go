package gocomp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {

	expected := Person{
		Name: "Bob",
		Age:  45,
	}

	actual := CreatePerson("Bob", 45)

	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	if diff := cmp.Diff(expected, actual, comparer); diff != "" {
		t.Error(diff)
	}
}
