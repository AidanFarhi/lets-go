package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
)

// you can use errors.New() to create a new error from a string
func DivMod(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return num / denom, num % denom, nil
}

// you can use fmt.Errorf() for formatting the error string
func DoubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d is not an even number", i)
	}
	return i * 2, nil
}

// including additional information with errors
type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

// dummy login function
func login(uid, pwd string) (string, error) {
	return "", errors.New("error logging in")
}

// using StatusErr to provide more details about what went wrong
func TryLogin(uid, pwd, file string) ([]byte, error) {
	token, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	return []byte(token), nil
}

func main() {

	x := 20
	y := 3
	rem, mod, err := DivMod(x, y)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(rem, mod)

	res, err := DoubleEven(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// sentinal errors
	data := []byte("this is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err = zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("not a zip file")
	}
}
