package main

import (
	"errors"
	"fmt"
)

func main() {
	// Map build key string and value int
	// make(mapType, n-elements) two arguments just for maps (no 3rd argument as for slice)
	maps := make(map[string]int)
	maps["hey"] = 42
	// short way to create map
	// maps := map[string]string{"hey": "world"}
	fmt.Println("map: ", maps)

	names := []string{"John", "Bob"}
	phones := []int{111, 222}
	userMap, err := getUserMap(names, phones)
	fmt.Println("user map:", userMap, err)

	john := user{name: "John", toDelete: true}
	// short way to create struct instance (should preserve an order)
	bob := user{"Bob", false}
	// maps can have struct as key and as value (comparable)
	// but cant have maps, functions, slices as key (not comparable)
	users := map[int]user{
		1: john,
		2: bob,
	}
	deleted1, err1 := deletion(users, 1)
	fmt.Println("user1 deletion:", deleted1, err1, users)
	deleted2, err2 := deletion(users, 2)
	fmt.Println("user2 deletion:", deleted2, err2, users)
	deleted3, err3 := deletion(users, 3)
	fmt.Println("user3 deletion:", deleted3, err3, users)

	persons := []string{"john", "bob", "ann"}
	fmt.Println("all persons:", addToMap(persons))

	films := map[string]string{"terminator": "action", "titanic": "drama"}
	// if there is no key in map, it will returns depends on type:
	// string > empty string; int > 0; bool > false ...
	fmt.Println("user does'nt exist:", films["empty"])

	// maps can have another map as a value (nested map)
	// rune is a type (literal is ''), represents single unicode symbol
	nestedMap := map[rune]map[string]int{
		// redundant syntax
		// "ann": map[string]int{"age": 23},
		// "john": map[string]int{"age": 32},
		// short syntax
		'a': {"age": 23},
		'j': {"age": 32},
	}
	fmt.Println("nested map:", nestedMap)
}

func addToMap(persons []string) map[string]int {
	// empty map can be created
	myMap := map[string]int{}
	for i, person := range persons {
		// empty map is getting fulfilled
		myMap[person] = i
	}
	return myMap
}

type user struct {
	name     string
	toDelete bool
}

func deletion(users map[int]user, id int) (deleted bool, err error) {
	// can be returned two values:
	// 2nd value (boolean - if exist) is optional
	user, ok := users[id]
	if !ok {
		return false, errors.New("there is no such a user")
	}
	// alternative (for above statement) syntax for if statement
	// if _, ok := users[id]; !ok { return false, errors.New("there is no such a user") }
	if !user.toDelete {
		return false, nil
	}
	// like slices, maps are also passed by reference into functions
	// if delete record from map inside the function,
	// original map will mutate
	delete(users, id)
	return true, nil
}

func getUserMap(names []string, phones []int) (map[string]int, error) {
	if len(names) != len(phones) {
		return nil, errors.New("lengths of arguments are not equal")
	}
	userMap := make(map[string]int)
	for i, name := range names {
		userMap[name] = phones[i]
	}
	return userMap, nil
}
