package setupteardown

import (
	"fmt"
	"os"
	"testing"
)

var data []int

func TestMain(m *testing.M) {
	fmt.Println("setting up data for tests")
	data = []int{1, 2, 3}
	fmt.Println("running tests")
	exitVal := m.Run()
	fmt.Println("cleaning up")
	os.Exit(exitVal)
}

func TestOne(t *testing.T) {
	fmt.Println("in TestOne data:", data)
}

func TestTwo(t *testing.T) {
	fmt.Println("in TestTwo data:", data)
}
