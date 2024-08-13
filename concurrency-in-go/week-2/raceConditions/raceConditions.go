package main

import (
	"fmt"
	"sync"
)

func counter(val *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		*val++
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	val := 0
	go counter(&val, &wg)
	go counter(&val, &wg)

	wg.Wait()
	fmt.Println(val)
}

/*
This race condition occurs because two goroutines are concurrently incrementing
the same variable 'val' without proper synchronization. The operation *val++ is not
atomic; it consists of multiple steps: reading the current value of 'val', incrementing
it, and writing the new value back to 'val'. When both goroutines execute these steps
at the same time, they can interfere with each other, leading to lost updates.
This causes the final value of 'val' to be less than expected, as some increments
are overwritten by others.
*/
