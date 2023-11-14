package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
)

func divMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
	Err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.Err
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {

	nums := [][]int{{20, 3}, {18, 2}, {55, 0}, {43, 5}}

	// handling errors while processing
	for _, pair := range nums {
		numerator := pair[0]
		denominator := pair[1]
		remainder, mod, err := divMod(numerator, denominator)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(remainder, mod)
		}
	}

	for _, pair := range nums {
		i := pair[0]
		result, err := doubleEven(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}

	// sentinal errors signal that processing cannot continue
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("not a zip file")
	}

	// wrapping and unwrapping errors
	err = fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}

	// checking returned errors
	err = fileChecker("does_not_exist.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file does not exist")
		}
	}

	// panic and recovering
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}
