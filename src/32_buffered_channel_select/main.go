package main

import (
	"fmt"
	"time"
)

func bufferChannelDemo() {
	c := make(chan int, 3)
	go func() {
		for {

			fmt.Println(<-c)
		}
	}()
	c <- 3
	c <- 4
	c <- 5
	c <- 6
}

func selectDemo() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	timeout := make(chan bool)

	go func() {
		time.Sleep(time.Millisecond)
		timeout <- true
	}()

	go func() {
		ch1 <- 10
		ch2 <- 20
	}()

	//可以使用select语句和default case来实现非阻塞
	for {
		select {
		case <-ch1:
			fmt.Println("receive data from ch1")
		case <-ch2:
			fmt.Println("receive data from ch2")
		case <-timeout:
			fmt.Println("timeout")
			return
		default:
			fmt.Println("no data")
		}
	}
}

//检查channel是否满了 非阻塞
func checkChannelFull(c chan int) {
	select {
	case c <- 2:
	default:
		fmt.Println("channel is full")
	}
}

func main() {

	//bufferChannelDemo()
	//selectDemo()
	c := make(chan int, 2)
	c <- 1
	c <- 2
	checkChannelFull(c)
	<-time.After(time.Millisecond)

}
