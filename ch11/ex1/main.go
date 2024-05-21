package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed english_rights.txt
var englishRights string

//go:embed gaelic_rights.txt
var gaelicRights string

//go:embed georgian_rights.txt
var georgianRights string

func main() {
	if len(os.Args) == 2 {
		var rights []string
		language := os.Args[1]
		switch language {
		case "english":
			rights = strings.Split(englishRights, "\n")
		case "gaelic":
			rights = strings.Split(gaelicRights, "\n")
		case "georgian":
			rights = strings.Split(georgianRights, "\n")
		}
		if len(rights) > 0 {
			for _, line := range rights {
				fmt.Println(line)
			}
		}
	} else {
		fmt.Println("usage: ./main <language>")
		os.Exit(1)
	}
	os.Exit(0)
}
