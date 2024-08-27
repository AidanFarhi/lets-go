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
	Err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.Err
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
			Err:     err,
		}
	}
	return []byte(token), nil
}

// incorrect usage of custom error
func GenerateErrorBroken(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

// correct usage
func GenerateErrorFixed(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	return nil
}

// wrapping errors
func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// merging multiple errors
func ValidatePerson(p Person) error {
	var errs []error
	if len(p.FirstName) == 0 {
		errs = append(errs, errors.New("field FirstName cannot be empty"))
	}
	if len(p.LastName) == 0 {
		errs = append(errs, errors.New("field LastName cannot be empty"))
	}
	if p.Age < 0 {
		errs = append(errs, errors.New("field Age cannot be negative"))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
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

	// For an interface to be considered nil, both the underlying type
	// and the underlying value must be nil. The underlying type of
	// these return values are not nil.
	err = GenerateErrorBroken(true)
	fmt.Println("GenerateErrorBroken(true) returns non-nil error:", err != nil)
	err = GenerateErrorBroken(false)
	fmt.Println("GenerateErrorBroken(false) returns non-nil error:", err != nil)

	err = GenerateErrorFixed(true)
	fmt.Println("GenerateErrorFixed(true) returns non-nil error:", err != nil)
	err = GenerateErrorFixed(false)
	fmt.Println("GenerateErrorFixed(false) returns non-nil error:", err != nil)

	err = fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}

	_, err = TryLogin("yolo", "pwd", "not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}

	p := Person{Age: -1}
	err = ValidatePerson(p)
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
}
