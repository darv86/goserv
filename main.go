package main

import "fmt"

func main() {
	var print = fmt.Printf

	// Instance creation
	myCar := Car{
		nameCar: "mazda",
		wheels:  4,
		// embedded struct can be created without having struct previously
		Engine: Engine{nameEngine: "rolls"},
	}

	// Instance creation alternative (separate fields assignment)
	// myCar := Car{}
	// myEngine := Engine{}
	// myEngine.nameEngine = "royal"

	// myCar.Engine = myEngine
	// myCar.nameCar = "mazda"
	// myCar.wheels = 4

	// Embedded struct
	// now can assign a value straight to a key of the Engine struct
	// myCar.nameEngine = "rolls royce"

	// Struct initialization
	// regular assignment typed value to the key
	// myCar.engine = myEngine
	// also can be done with anonymous struct
	// myCar.engine = struct{ name string }{name: "jaguar"}

	print("there is res1 - %v", myCar)
}

// Embedded struct
type Car struct {
	nameCar string
	wheels  int
	// struct Car will have all fields of Engine struct
	Engine
}

type Engine struct {
	nameEngine string
}

// Struct initialization
// type Car struct {
// 	name   string
// 	wheels int
// 	// regular field typing
// 	// engine Engine
// 	// also can be done with anonymous struct as well (in this case contracts needs to be valid)
// 	engine struct{ name string }
// }
// type Engine struct {
// 	name string
// }
