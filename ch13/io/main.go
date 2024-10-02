package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	g, err := gzip.NewReader(f)
	if err != nil {
		return nil, nil, err
	}
	return g, func() {
		f.Close()
		g.Close()
	}, nil
}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func main() {

	// simple reader example
	f, err := os.Open("words.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	letterCounts, err := countLetters(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(letterCounts)

	// gzip reader example
	gzipReader, closer, err := buildGZipReader("compressed_words.txt.gz")
	defer closer()
	letterCounts, err = countLetters(gzipReader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(letterCounts)
}
