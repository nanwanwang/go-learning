package main

import (
	"fmt"
	"sync"
)

var x int

var wg sync.WaitGroup

var lock sync.Mutex

var rwlock sync.RWMutex

func add() {

	defer wg.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}


func read(){
   

}


func write(){

	
}



func main() {

	wg.Add(2)
	go add()
	go add()

	wg.Wait()

	fmt.Println(x)

}
