package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	mutex sync.Mutex
)

func main() {
	for i := 0; i < 10000; i++ {
		go increment()
	}
	time.Sleep(time.Second * 1)
	fmt.Println("Count:", count)
}

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	count++
}
