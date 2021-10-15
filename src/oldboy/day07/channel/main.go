package main

import (
	"fmt"
	"sync"

)

// var c chan int

var wg sync.WaitGroup

func chanNoBuffer() {
	var c chan int
	c = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		num := <-c
		fmt.Printf("获取到值%d\n", num)
	}()

	c <- 10

	defer close(c)

}

func send(chan1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		chan1 <- i
	}
	close(chan1)
}


//单向通道  chan1 发送  chan2 只能接收
func receive(chan1 <-chan int, chan2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-chan1
		if !ok {
			break
		}
		chan2 <- x * x
	}
	close(chan2)

}

func main() {

	// c = make(chan int,1)

	// c <- 2
	// c <- 3

	// num:=<-c

	// fmt.Println(num)

	// fmt.Println(c)

	// chanNoBuffer()

	var chan1 chan int = make(chan int, 100)
	var chan2 chan int = make(chan int, 100)
	wg.Add(2)

	go send(chan1)
	go receive(chan1, chan2)

	wg.Wait()
	for res := range chan2 {
		fmt.Println(res)
	}

}
