package main

import (
	"fmt"
)

func main() {

	logger := getLogger(doFormat)
	logger("1st string ", "2nd string")

	fmt.Println("defer call:", def())
}

func def() (res int) {
	res = 5
	// Defer keyword postpone func call,
	// after value (res) in return statement is ready
	defer func() { res *= 2 }()
	// its convenient, coz there is no need of calling
	// func (in this case, anonymous) inside every return statement
	if res > 4 {
		return
	}
	res += 3
	return res
}

// High order function
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(str1, str2 string) {
		fmt.Println(formatter(str1, str2))
	}
}

// In Go (the same in js) all functions are 1st class functions:
// they can be passed as a argument to another function
// or assigned to a variable
// or returned as a value
func doFormat(str1, str2 string) string {
	return str1 + str2
}
