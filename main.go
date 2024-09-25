package main

import "fmt"

func main() {
	// fmt.Print("hello")
	// fmt.Print("world")
	// fmt.Println("world")
	// fmt.Println("hello")

	var print = fmt.Printf

	// const confName = "Conf name"
	// tickets := 20.123 // var tickets = 20
	// ticketsInt := int(tickets)
	// ticketsStr := fmt.Sprintf("ticketsStr - %.1f", tickets)

	// print(
	// 	"here is %v\ntickets %.2f\nticketsInt - %v\n%v",
	// 	confName,
	// 	tickets,
	// 	ticketsInt,
	// 	ticketsStr,
	// )

	// if age := 17; age >= 18 {
	// 	print("you can, you are %v", age)
	// } else if age < 18 {
	// 	print("you can't, you are %v", age)
	// } else {
	// 	print("i need your age")
	// }

	// res := adder(2, 3, "my string\n")

	// res := 10
	// const now int = 1
	// res = plus(res, now)

	// res, res2 := multipleReturn("string1", "string2")
	res2, res1 := multipleReturn(2, -3)

	// print("there is %v", res)
	print("there is res1-%v and res2-%v", res1, res2)
}

// func multipleReturn(s1, s2 int) (str2, str1 int) {
// 	str1 = s1
// 	str2 = s2
// 	if str1 < 0 {
// 		str1 = 0
// 	}
// 	if str2 < 0 {
// 		str2 = 0
// 	}
// 	return // you can omit explicit values in return statement, if all values should be returned
// }

func multipleReturn(s1, s2 int) (int, int) {
	str1 := s1
	str2 := s2
	if str1 < 0 {
		str1 = 0
	}
	if str2 < 0 {
		str2 = 0
	}
	return str2, str1 // the same as above
}

// func plus(a, b int) int {
// 	return a + b
// }

// func adder(a int, b int, str) int {
// 	return a + b
// }
// func adder(a, b int, str string) int {
// 	print(str)
// 	return a + b
// }
