package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      = 0
	lock   sync.Mutex
	wg     sync.WaitGroup
	rwLock sync.RWMutex
)

func write() {
	defer wg.Done()
	lock.Lock()
	x = x + 1
	time.Sleep(time.Microsecond * 5)
	lock.Unlock()

}

func read() {
	defer wg.Done()
	lock.Lock()
	fmt.Println(x)
	lock.Unlock()

}

func main() {

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	time.Sleep(time.Second)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start))

}
