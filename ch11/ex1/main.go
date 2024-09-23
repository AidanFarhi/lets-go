package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed text/english_rights.txt
var english string

//go:embed text/greek_rights.txt
var greek string

//go:embed text/hmong_rights.txt
var hmong string

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: rights <lang>")
		os.Exit(0)
	}
	switch os.Args[1] {
	case "eng":
		fmt.Println(english)
	case "gre":
		fmt.Println(greek)
	case "hmo":
		fmt.Println(hmong)
	default:
		fmt.Println("language not found")
	}
}
