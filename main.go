package main

import (
	"fmt"
	"github.com/darv86/goserv/reverser"
)

func main() {
	textR := reverser.Reverse("my text")
	fmt.Println("reversed text:", textR)
}
