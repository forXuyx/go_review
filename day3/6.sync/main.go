package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
