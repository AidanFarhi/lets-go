package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileLen(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	length := 0
	chunk := make([]byte, 2048)
	for {
		count, err := f.Read(chunk)
		length += count
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}
	return length, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	fileName := os.Args[1]
	length, err := fileLen(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileName, "length:", length)
}
