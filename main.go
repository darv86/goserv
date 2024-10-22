package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	list := map[string]int{"bob": 1, "john": 2, "ann": 3}
	// Mutex initialization
	mtx := &sync.Mutex{}
	go adder(list, mtx)
	go multiplier(list, mtx)
	time.Sleep(time.Second)
	fmt.Println("final list after Mutex:", list)

	rMtx := &sync.RWMutex{}
	data := &[]int{}
	go reader1(data, rMtx)
	go reader2(data, rMtx)
	go writer(data, rMtx)
	// infinite loop
	// select {}

	a := []int{1}
	func(data []int) {
		// slice changes are visible outside this func
		// data[0] = 2
		// slice changes are not visible outside this func
		// and append creates copy of a slice (data)
		data = append(data, 2)
		fmt.Println("final data:", data)
	}(a)
	fmt.Println("final data:", a)
}

func writer(list *[]int, rMtx *sync.RWMutex) {
	// RLock method blocks data (list) only for writing
	// all functions (reader1, reader2, writer),
	// which have mutex (rMtx) with RLock method
	// can read data (list) at the same time,
	// but write at the same time can only one
	rMtx.RLock()
	defer rMtx.RUnlock()
	for i := 0; ; i++ {
		// If the underlying array’s capacity is exceeded,
		// (e.g.: append exceeds underlying array’s capacity)
		// append creates a new array,
		// and the slice within the function
		// won’t reflect this change outside the function
		// that's why pointer is used here:
		// this func changes slice (list),
		// using reference on the original array
		*list = append(*list, i)
		fmt.Println("from writer:", (*list)[i])
		time.Sleep(500 * time.Millisecond)
	}

}

func reader1(list *[]int, rMtx *sync.RWMutex) {
	rMtx.RLock()
	defer rMtx.RUnlock()
	for {
		fmt.Println("from reader1:", *list)
		time.Sleep(500 * time.Millisecond)
	}
}

func reader2(list *[]int, rMtx *sync.RWMutex) {
	rMtx.RLock()
	defer rMtx.RUnlock()
	for {
		fmt.Println("from reader2:", *list)
		time.Sleep(500 * time.Millisecond)
	}
}

func adder(list map[string]int, mtx *sync.Mutex) {
	// Lock method blocks data (list) from reading and writing,
	// while this function (adder) executes
	mtx.Lock()
	// Unlock method unblock data, after this func is done
	defer mtx.Unlock()
	for name, id := range list {
		list[name] = id + 2
		fmt.Println("from adder name:", name, "--- id:", list[name])
	}
}

func multiplier(list map[string]int, mtx *sync.Mutex) {
	mtx.Lock()
	defer mtx.Unlock()
	for name, id := range list {
		list[name] = id * 2
		fmt.Println("from multiplier name:", name, "--- id:", list[name])
	}
}
