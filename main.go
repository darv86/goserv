package main

import "fmt"

func main() {
	var print = fmt.Printf

	myCar := Car{
		nameCar: "mazda",
		wheels:  4,
		engine:  struct{ nameEngine string }{nameEngine: "rolls"},
	}

	print("there is result - %v, %s", myCar, myCar.getInfo(";"))
}

type Car struct {
	nameCar string
	wheels  int
	engine  struct{ nameEngine string }
}

// Method for struct Cat
func (c Car) getInfo(separator string) string {
	return fmt.Sprintf(
		"name: %s%s engine: %s",
		c.nameCar,
		separator,
		c.engine.nameEngine,
	)
}
