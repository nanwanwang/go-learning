package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func basicSelect(ch1 chan<- int, ch2 <-chan int) {
	for i := 0; i < 10; i++ {
		select {
		case ch1 <- 1:
			fmt.Println("接收到数据")
		case v := <-ch2:
			fmt.Printf("取到数据%d\n", v)
		default:
			fmt.Println("default")

		}
	}
}

func main() {

	// ch1:=make(chan int,2)
	// ch2:=make(chan int,2)
	// ch2<-2
	// ch2<-10
	// basicSelect(ch1,ch2)

	// var ch = make(chan int,1)

	// //wg.Add(1)
	// go func(c chan int){
	// 	//defer wg.Done()
	// 	c<- 100
	// }(ch)

	// //wg.Wait()

	// for{
	// 	select{
	// 	case <- ch:
	// 		fmt.Println("ok")
	// 	default:
	// 		fmt.Println("default")
	// 	}

	// 	time.Sleep(time.Second)
	// }

	ch := make(chan int)


	go func(c chan int) {
		time.Sleep(time.Second*1)
		ch <- 1
	}(ch)

	for{
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-time.After(2 * time.Second):
			fmt.Println("no case ok")
		}
	}
	
	//time.Sleep(time.Second * 10)

}
