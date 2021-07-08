package main

import (
	"fmt"
	"time"
)

func channelDemo() {
	c := make(chan int)
	go func() {
		for{
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 10

	time.Sleep(time.Millisecond)

}

func main() {
	channelDemo()
	//fmt.Println(1)
}
