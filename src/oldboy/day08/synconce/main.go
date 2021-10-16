package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

var m1 sync.Map

var lock sync.Mutex

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {

	var wg = sync.WaitGroup{}
	// for i:=0;i<20;i++{
	// 	wg.Add(1)
	// 	go func(n int){
	// 		key:=strconv.Itoa(n)
	// 		lock.Lock()
	// 		set(key,n)
	// 		lock.Unlock()
	// 		fmt.Printf("k:=%v,v:=%v\n",key,get(key))
	// 		defer wg.Done()
	// 	}(i)
	// }

	// wg.Wait()

	//sync.Map

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m1.Store(key, n)
			if value, ok := m1.Load(key); ok {
				fmt.Printf("k:=%v,v:=%v\n", key, value)
			}

			defer wg.Done()

		}(i)
	}

	wg.Wait()

   atomic.lo

}
