package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel creation
	// use channel to receive value from goroutine
	ch := make(chan int)

	// this is nill channel - causes block while reading or writing
	// var cn chan int
	// fmt.Println(<-cn)

	// Goroutine is more lightweight than a traditional thread
	// go keyword runs func concurrently
	// with other similar functions in there own goroutines
	// block #1 and block #2 run concurrently
	go gor(ch, 3)
	// when channel receive a value from goroutine,
	// this value can be extracted, using <- syntax
	// and after can be assigned to the variable (gv)
	// gv := <-ch
	// block #2
	// there are 3 extractions from channel (ch)
	// and this is a limit for loop in goroutine (gor)
	// even if loop has more then 3 iterations,
	// func (gor) stops after 4th iteration
	fmt.Println("value from goroutine", <-ch)
	fmt.Println("value from goroutine", <-ch)
	fmt.Println("value from goroutine", <-ch)

	num := 3
	c, _ := getDBsChannel(num)
	waitForDBs(num, c)

	// Buffered channel with optional argument (buffer size)
	ch1 := make(chan int, num)
	// if size of a buffer is equal to quantity of received items
	// you can use functions in one single (in this case, main) goroutine
	// (functions without go keyword)
	sender(num, ch1)
	receiver(num, ch1)

	ch2 := make(chan int)
	go sendReports(3, ch2)
	count := countReports(ch2)
	fmt.Println("reports sent:", count)

	fibSlice := concurrentFib(10)
	fmt.Println("fibonacci:", fibSlice)

	chInt := make(chan int)
	chStr := make(chan string)
	go intGen([]int{1, 2, 3}, chInt)
	go strGen([]string{"one", "two", "three"}, chStr)
	chReader(chInt, chStr)

	// Tick sends value (current time) to returned channel (timeCh)
	// every time a duration (Tick argument) elapses
	// tickCh := time.Tick(3 * time.Second)
	// for v := range tickCh {
	// 	fmt.Println("from tick channel:", v)
	// }

	// After sends the current time on the returned channel
	// after delay (After argument)
	// afterCh := time.After(5 * time.Second)
	// fmt.Println("from after channel", <-afterCh)

	// Sleep stops the current goroutine
	// (in this case, func main is the current goroutine)
	// for the duration (Sleep argument)
	// fmt.Println("before sleep", time.Now())
	// time.Sleep(3 * time.Second)
	// fmt.Println("after sleep", time.Now())

	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	logChan := make(chan string)
	go saveBackups(snapshotTicker, saveAfter, logChan)
	actualLogs := []string{}
	for actualLog := range logChan {
		fmt.Println("from main:", actualLog)
		actualLogs = append(actualLogs, actualLog)
	}
	fmt.Println(actualLogs)
}

// func watcher(ch <-chan time.Time) {
// 	select {
// 	case v := <-ch:
// 		fmt.Println("from watcher:", v)
// 	}
// }

// num <- chan int (argument syntax only for read from a channel)
// num chan <- int (argument syntax only for write in a channel)
func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			// if !ok {
			// 	break
			// }
			takeSnapshot(logChan)
		case <-saveAfter:
			// if !ok {
			// 	break
			// }
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}
func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}
func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

func chReader(chInt chan int, chStr chan string) {
	count := []int{}
	for {
		if len(count) == 2 {
			fmt.Println("count is full:", count)
			break
		}
		// Select is like switch in js, but only for channels
		// select watches at the same time for a few channels,
		// which one first will assign value to a variable (v),
		// this case works first
		// if channel get closed, second variable (ok) becomes false
		select {
		case v, ok := <-chInt:
			if !ok {
				count = append(count, 1)
			}
			fmt.Println("from chInt:", v, ok)
		case v, ok := <-chStr:
			if !ok {
				count = append(count, 2)
			}
			fmt.Println("from chStr", v, ok)
		// prevent select from blocking
		// default does operation, if nothing to receive from channel
		// and channel doesn't close
		default:
			fmt.Println("from default")
		}
	}
}

func intGen(num []int, ch chan int) {
	for _, v := range num {
		ch <- v
	}
	close(ch)
}

func strGen(num []string, ch chan string) {
	for _, v := range num {
		ch <- v
	}
	close(ch)
}

func concurrentFib(n int) []int {
	ch := make(chan int)
	res := []int{}
	go fibonacci(n, ch)
	// you can loop over channel, like slice or map
	// if channel get closed, loop breaks
	// iteration get blocked, till value received from channel
	for v := range ch {
		res = append(res, v)
	}
	return res
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func countReports(numSentCh chan int) int {
	count := 0
	for {
		v, ok := <-numSentCh
		if !ok {
			break
		}
		count += v
	}
	return count
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	// read from a closed channel, returns 0 value (chan int)
	// send to a closed channel, panics
	close(ch)
}

func sender(num int, ch chan int) {
	for i := range num {
		ch <- i
	}
	// when operation with channel is done,
	// you can close the channel
	// do this from the sender side
	close(ch)
}

func receiver(num int, ch chan int) {
	for range num + 1 {
		// ok (bool) is second optional variable
		// it indicates, if channel is closed
		// but if channel is buffered,
		// ok will be false, if channel is closed and empty
		v, ok := <-ch
		fmt.Println("from receiver:", v, ok)
	}

}

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for range numDBs {
		// this receiver blocks runtime,
		// till new struct{} will be passed to this channel
		// in the goroutine (anonymous func inside getDBsChannel)
		<-dbChan
	}
}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			// passing struct{} to the channel
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()
	return ch, &count
}

func gor(ch chan int, n int) {
	for i := range n + 1 {
		fmt.Println("from loop", i)
		// goroutine passes value to channel
		// and indicates about completion
		// (almost like return does, but multiple times)
		ch <- i
	}
	// if a sender (ch <-) sends to channel the value (i),
	// but there is no receiver (loop iterations n+1 > receivers in func main)
	// to receive particular value,
	// this sender blocks code execution on the current line
	fmt.Println("from loop: end")
}
