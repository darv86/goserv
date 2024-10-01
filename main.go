package main

import "fmt"

func main() {
	res1 := loop(5)
	fmt.Println("result from loop:", res1)

	res2 := while(5)
	fmt.Println("result from while:", res2)

	rem := 24 % 7
	fmt.Println("rem is:", rem)
}

func loop(end int) int {
	count := 0
	// i < end (condition section) works before iteration
	// i++ (after section) works after iteration
	for i := 0; i < end; i++ {
		count = i
		// continue and break keywords as in js
	}
	return count
}

func while(end int) int {
	count := 0
	// While loop in go looks like for loop,
	// but without Initial and After section
	for count < end {
		count++
	}
	return count
}
