package main

import "fmt"

func main() {
	print := fmt.Printf
	// println := fmt.Println

	myCar := Car{
		nameCar: "mazda",
		wheels:  4,
		engine:  struct{ nameEngine string }{nameEngine: "rolls"},
	}

	print(
		"there is result - %v, %s, %v, %v",
		myCar,
		myCar.getInfo(";"),
		// getInfoLen accepted myCar as argument, but Informer in function declaration
		getInfoLen(myCar),
		myCar.getNameLen(2),
	)
}

// Interface - collection of a methods signatures
type Informer interface {
	getInfo(string) string
}

type NameCounter interface {
	// argument can be named ("offset")
	// returned value also can be named ("length")
	// arguments and returned values can have own name,
	// but should not have the same name as original methods
	getNameLen(offset int) (length int)
}

// can accept Car instance, coz Car implicitly implements (has getInfo methods) Informer interface
func getInfoLen(i Informer) int {
	// Type assertion
	// inf - Car instance
	// ok - true if "i" has underlying type of Car
	// inf, ok := i.(Car)
	// fmt.Println(inf, ok)
	// Type switches
	switch v := i.(type) {
	// actually v == i
	case Car:
		// will print type of argument (Car)
		fmt.Printf("from switch Car %v\n", v)
	case Plane:
		// do almost the same (main.Plane)
		fmt.Printf("from switch Plane %T\n", i)
	default:
		fmt.Printf("from switch default %v\n", i)
	}
	return len(i.getInfo(""))
}

type Plane struct{}

func (p Plane) getInfo(separator string) string {
	return "from plane"
}

type Car struct {
	nameCar string
	wheels  int
	engine  struct{ nameEngine string }
}

func (c Car) getInfo(separator string) string {
	return fmt.Sprintf(
		"name: %s%s engine: %s",
		c.nameCar,
		separator,
		c.engine.nameEngine,
	)
}

// now Car implements second interface NameCounter
func (c Car) getNameLen(offset int) (l int) {
	l = offset + len(c.nameCar)
	return l
}
