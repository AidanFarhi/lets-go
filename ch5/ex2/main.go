package main

import (
	"fmt"
	"os"
)

func fileLen(fileName string) (int, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return 0, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return int(fileInfo.Size()), nil
}

func main() {
	file1 := "words.txt"
	file2 := "i_dont_exist.txt"
	file1Len, err := fileLen(file1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(file1Len)
	file2Len, err := fileLen(file2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(file2Len)
}
