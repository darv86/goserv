package main

import (
	"fmt"
)

func main() {
	arrInt := []int{1, 2, 3, 4}
	slc1, slc2 := splitSlice(arrInt)
	fmt.Println("generics", slc1, slc2)

	arrStr := []string{"one", "two", "three"}
	fmt.Println("last element:", getLast(arrStr))
	fmt.Println("last zero element:", getLast([]int{}))

	user := &user{name: "bob"}
	admin := &admin{name: "ann"}
	fmt.Println("converter result (user):", converter(user))
	fmt.Println("converter result (admin):", converter(admin))

	fmt.Println("first element:", getFirst(arrStr))
	// Custom type (myInt) creation with underlying type of int
	type myInt int
	myArrInt := []myInt{1, 2, 3, 4}
	// error if argument for gerFirst is custom type (myInt)
	// or there is no ~ syntax inside interface (listItem)
	fmt.Println("first element:", getFirst(myArrInt))

	// vault implements store interface,
	// which has type parameter (si),
	// this allows to create two vaults (string, int),
	// then use the same methods (getVault, addToVault),
	// and func (vaultLogger) on these vaults
	vaultInt := &vault[int]{items: []int{}}
	vaultInt.addToVault(1, 2, 3)
	vaultLogger(vaultInt)
	vaultStr := &vault[string]{items: []string{}}
	vaultStr.addToVault("one", "two", "three")
	vaultLogger(vaultStr)
	fmt.Println("data from both vaults: ", vaultStr.getVault(), vaultInt.getVault())
}

type si interface {
	string | int
}

// Type parameter in interface
type store[P si] interface {
	// type parameter can be passed
	// to the specific method of an instance,
	// which implements this interface
	getVault() []P
	addToVault(...P)
}

// store interface is needed to have one func,
// which can accept two types (string, int = si) of vaults;
// methods in go can't have generic type and
// custom interface with constraints as a type for parameters:
// func (mo *mObj) addToVault(item unit)
// but functions can:
func vaultLogger[P si](vault store[P]) {
	fmt.Printf("data from %T: %v\n", vault, vault.getVault())
}

// Type parameter in struct;
// type parameter (P si) further
// can be used for the specific method
type vault[P si] struct {
	items []P
}

// value receiver method works fine with pointer (vaultInt),
// because Go automatically dereferences the pointer;
// Type parameters in the method of the vault instance
func (v vault[P]) getVault() []P {
	return v.items
}

// use pointer receiver method,
// if you need to modify the struct (vaultInt);
// Type parameters in the method of the vault instance
func (v *vault[P]) addToVault(item ...P) {
	v.items = append(v.items, item...)
}

// -----------------------------

func getFirst[T listItem](arr []T) T {
	element := arr[0]
	return element
}

type listItem interface {
	// Union type constraint
	// generic (listItem) func (getFirst)
	// can accept only array of strings or ints
	// string | int
	// ~ syntax says that any type,
	// that has underlying type of string,
	// can be type of listItem
	// if there is no ~, strict type of string or int should be in use
	~string | ~int
}

// this generic (T) has constraint (stringer) via interface:
// incoming argument (src) can be any struct,
// that implements interface (stringer)
func converter[T stringer](src T) string {
	return fmt.Sprint("-> ", src.getName())
}

type stringer interface {
	getName() string
}

type user struct {
	name string
}
type admin struct {
	name string
}

func (u *user) getName() string {
	return u.name
}
func (a *admin) getName() string {
	return a.name
}

// -------------------------------

func getLast[T any](store []T) T {
	if len(store) == 0 {
		// to return zero value of the specific type
		// use this syntax
		var zeroVal T
		return zeroVal
	}
	item := store[len(store)-1]
	return item
}

// -------------------------------

func splitSlice[T any](slice []T) ([]T, []T) {
	mid := len(slice) / 2
	fmt.Println("mid", mid)
	return slice[:mid], slice[mid:]
}
