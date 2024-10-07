package main

import "fmt"

func main() {
	// Array in go has fixed length
	arr := [4]int{2, 4, 6, 3}
	fmt.Println("result from array:", arr)
	// Slice is a view for array
	// in most cases slice uses to have a deal with array
	// slices store value by reference from array, it means,
	// if function has access to slice, function can modify array
	slice1 := arr[:]
	fmt.Println("result from slice1:", slice1)
	slice2 := arr[1:]
	fmt.Println("result from slice2:", slice2)
	slice3 := arr[:3]
	fmt.Println("result from slice3:", slice3)
	slice4 := arr[1:3]
	fmt.Println("result from slice4:", slice4)

	// slice literal (doesn't use length value inside array literal[], but array does)
	arrStr := []string{"a", "b", "c"}
	fmt.Printf("result from arrStr: %T, %T\n", arrStr, arr)

	// slice can be created over another slice
	arrStrSlice := arrStr[:2]
	fmt.Println("result from arrStrSlice:", arrStrSlice)

	// another way to create a slice
	// make(type, length, capacity)
	sliceMake := make([]int, 3, 6)
	sliceMakeLen := len(sliceMake)
	sliceMakeCap := cap(sliceMake)
	fmt.Printf("result from sliceMake: %v; length: %v; capacity: %v\n", sliceMake, sliceMakeLen, sliceMakeCap)

	arrayForSlice := [6]int{1, 2, 3, 4, 5, 6}
	sliceFromArray := arrayForSlice[:2]
	fmt.Printf(
		"array: %v; length: %v; cap: %v\nslice: %v, length: %v; cap: %v\n",
		arrayForSlice,
		len(arrayForSlice),
		cap(arrayForSlice),
		sliceFromArray,
		len(sliceFromArray),
		// The capacity of a slice is the number of elements in the underlying array,
		// counting from the first element in the slice
		cap(sliceFromArray),
	)

	// sum receives arguments as many as needed
	fmt.Println("total: ", sum(2, 4, 1))

	sliceToRest := []int{2, 5, 1}
	// Rest operator in go
	fmt.Println("total with rest: ", sum(sliceToRest...))

	costs := []costPerDay{
		{0, 4.0},
		{1, 2.1},
		{5, 2.5},
		{1, 3.1},
	}
	fmt.Println("total costs:", getCostByDay(costs))

	fmt.Println("matrix:", createSlice2d(10, 10))

	// Range can be used instead of:
	// for i:=0; i<len(arrFroRange); i++
	arrForRange := []int{1, 2, 3, 4}
	for i, elem := range arrForRange {
		fmt.Println("index and element:", i, elem)
	}
}

// Slice of slices type
func createSlice2d(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = i * j
		}
	}
	return matrix
}

// Variadic arguments (like rest operator in js)
func sum(numbers ...int) int {
	total := 0
	// numbers is literally a slice
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

type costPerDay struct {
	day  int
	cost float64
}

func getCostByDay(costs []costPerDay) []float64 {
	costsOnly := []float64{}
	days := 0
	for i := 0; i < len(costs); i++ {
		if costs[i].day > days {
			days = costs[i].day
		}
	}
	for i := 0; i <= days; i++ {
		// Append adds element to the end of a slice (like push in js)
		// and auto increase length of a new slice
		// reassign new slice to a current (slice) variable
		// try to avoid to assign new slice to another (slice) variable
		costsOnly = append(costsOnly, 0.0)
	}
	for i := 0; i < len(costsOnly); i++ {
		for j := 0; j < len(costs); j++ {
			if i == costs[j].day {
				costsOnly[i] += costs[j].cost
			}
		}
	}
	return costsOnly
}
