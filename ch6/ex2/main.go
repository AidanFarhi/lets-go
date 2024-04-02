package main

import "fmt"

func UpdateSlice(strSlice []string, str string) {
	strSlice[len(strSlice)-1] = str
	fmt.Println("In UpdateSlice: ", strSlice)
}

func GrowSlice(strSlice []string, str string) {
	strSlice = append(strSlice, str)
	fmt.Println("In GrowSlice: ", strSlice)
}

func main() {
	strings := []string{"one", "two", "three"}
	fmt.Println("Before UpdateSlice: ", strings)
	UpdateSlice(strings, "four")
	fmt.Println("After UpdateSlice: ", strings)
	fmt.Println("Before GrowSlice: ", strings)
	GrowSlice(strings, "five")
	fmt.Println("After GrowSlice: ", strings)
}
