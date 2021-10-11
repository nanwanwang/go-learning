package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("hello", i)
	}

	n := 0
	for n < 10 {
		fmt.Println(n)
		n++
	}

	for {
		fmt.Println("hello ")
		time.Sleep(1000)
	}

}
