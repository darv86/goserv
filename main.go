package main

import (
	"fmt"
	"strings"
)

func main() {
	x := 2
	// Pointer (p) stores the address of x
	var p *int = &x
	// p := &x // the same above
	// value of x, changed by the address,
	// so pointer (p) has updated value of x (3)
	x = 3
	// pointer stores the address of x
	fmt.Printf("pointer stores address: %v\n", p)
	// syntax * shows the value behind x address
	fmt.Printf("pointer stores value: %v\n", *p)
	*p = 4
	// updates the value for x
	fmt.Printf("x stores updated (by *p) value: %v\n", x)

	user := user{name: "bob"}
	str := "some string"
	// passed value with pointer type as argument
	printName(&user, &str)
	// printName func mutates struct (user)
	fmt.Println("user's changed name:", user.name)
	fmt.Println("and changed str:", str)
}

type user struct {
	name string
}

// after * syntax, argument (u) stores
// the address of user instance
// so if we want to mutate instance inside a method
// and keep changes outside this method,
// we should use pointer type as an argument for the method
func printName(u *user, str *string) {
	// if pointer points to nothing, it returns nil
	// so check if nil before dereferencing
	if u == nil || str == nil {
		return
	}
	fmt.Printf("user name is %v\n", u.name)
	// code above does the same as code below under the hood
	fmt.Printf("user name is %v\n", (*u).name)
	// and now we can mutate struct (user) inside this func
	u.name = "john"
	// 2nd * converts from pointer (*str) to string value for ReplaceAll,
	// replaces string and returns result
	// 1st * converts, converted value reassigns with new returned value
	*str = strings.ReplaceAll(*str, "string", "text")
}
