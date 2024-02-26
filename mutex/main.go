package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mutex sync.Mutex
var waitGroup sync.WaitGroup

func addValueToCounter(counter *int, valueToAdd int) {
	mutex.Lock()

	value := counter

	runtime.Gosched()

	(*value) += valueToAdd

	counter = value

	mutex.Unlock()

	waitGroup.Done()
}

func main() {
	const maxGoRoutines = 10
	const valueToAdd = 2
	
	counter := 0

	waitGroup.Add(maxGoRoutines)

	for iterator := 0; iterator < maxGoRoutines; iterator++ {
		go addValueToCounter(&counter, valueToAdd)
	}

	waitGroup.Wait()
	fmt.Println(counter)
}