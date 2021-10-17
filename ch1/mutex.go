package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go counter(&wg)
	}

	wg.Wait()
	fmt.Println(count)
}

func counter(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	count++
	mu.Unlock()
}
