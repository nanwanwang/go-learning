package main

import (
	"fmt"
	"time"
)

func getData(i int, c chan int) {
	//for {
	//	//n := <-c
	//	if n, ok := <-c; ok {
	//		fmt.Printf("get data %d from channel %d\n", n, i)
	//	} else {
	//		break
	//	}
	//}
	//range从通道中取值,如果channel关闭,自动终止本次循环
	for n := range c {
		fmt.Printf("get data %d from channel %d\n", n, i)
	}
}

func createChannel(i int) chan<- int {
	c := make(chan int)
	go getData(i, c)
	return c
}

func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createChannel(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- i + 100
	}

	for i := 0; i < 10; i++ {
		close(channels[i])
	}

	//channels[0] <- 22
	//c := createChannel()
	//go getData(c)
	//c <- 10
	//c <- 20
	time.Sleep(time.Second)

}

func main() {
	channelDemo()
	//fmt.Println(1)
}
