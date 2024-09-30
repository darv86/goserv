package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// if there is no need to use one of two returned values,
	// use "_" to ignore that value
	num, _ := strToInt("32b")
	err := myError{nameType: "myError", nameError: "this my error"}
	res, errDiv := div(2, 0)

	fmt.Printf(
		"there is result - %v, %T\n",
		num,
		num,
	)
	fmt.Println("my custom error message:", err)
	fmt.Println("result is:", res, errDiv)
}

func div(a, b int) (int, error) {
	if a == 0 || b == 0 {
		// New error (from standard lib)
		// creates error once
		return 0, errors.New("cant divide by zero")
	}
	return a / b, nil
}

func strToInt(str string) (int, error) {
	// strconv.Atoi - converts string to int
	num, err := strconv.Atoi(str)
	// Error handling (in go error is just regular value, not special type)
	// if causes error, err will have string with details,
	// if not, err == nil
	if err != nil {
		fmt.Println("error from if:", err)
		return 0, err
	}
	return num, err
}

// Custom error
// to use my own error with specific error message,
// myError type should implements global error interface:
// struct should have Error method, that returns string
type myError struct {
	nameError string
	nameType  string
}

// this method will call, when myError struct produces new instance
func (e myError) Error() string {
	// will return custom error message
	return e.nameError
}
