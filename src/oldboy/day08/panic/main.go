package main

import (
	"fmt"
	"time"
)

func G() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("c")
	}()

	go F()

	time.Sleep(time.Second)

}

func F() {
	defer func() {
		// if err := recover(); err != nil {
		// 	fmt.Println("捕获异常:", err)
		// }
		fmt.Println("b")
	}()

	panic("a")
}

func main() {
	G()
}
