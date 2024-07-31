package main

import "fmt"

func UpdateSlice(sl []string, s string) {
	sl[len(sl)-1] = s
	fmt.Println("in UpdateSlice:", sl)
}

func GrowSlice(sl []string, s string) {
	sl = append(sl, s)
	fmt.Println("in GrowSlice:", sl)
}

func main() {

	names := []string{"bob", "joe", "john", "moses"}

	fmt.Println("before UpdateSlice:", names)
	UpdateSlice(names, "abraham")
	fmt.Println("after UpdateSlice:", names)

	fmt.Println("before GrowSlice:", names)
	GrowSlice(names, "joseph")
	fmt.Println("after GrowSlice:", names)
}
