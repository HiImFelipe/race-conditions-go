package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var waitGroup sync.WaitGroup

func addValueToCounter(counter *int64, valueToAdd int64) {
	atomic.AddInt64(counter, valueToAdd)

	fmt.Println("Current counter value:", *counter)

	runtime.Gosched()

	waitGroup.Done()
}

func main() {
	const maxGoRoutines = 10
	const valueToAdd = 2
	var counter int64 = 0;

	waitGroup.Add(maxGoRoutines)

	for iterator := 0; iterator < maxGoRoutines; iterator++ {
		go addValueToCounter(&counter, valueToAdd)
	}

	waitGroup.Wait()
	fmt.Println("Final value:", counter)
}