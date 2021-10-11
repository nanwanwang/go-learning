package main

import (
	"fmt"
	"sync"
)

type safeInt struct {
	a    int
	lock sync.Mutex
}

func (si *safeInt) increase() {
	si.lock.Lock()
	defer si.lock.Unlock()
	si.a++
}

func (si *safeInt) get() int {
	si.lock.Lock()
	defer si.lock.Unlock()
	return si.a
}

func main() {
	var si safeInt
	var wg sync.WaitGroup
	si.increase()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			si.increase()
		}()
	}
	wg.Wait()
	//time.Sleep(time.Millisecond)
	fmt.Println(si.get())

	//var a int
	//a++
	//go func() {
	//	a++
	//}()
	//time.Sleep(time.Millisecond)
	//fmt.Println(a)
}
