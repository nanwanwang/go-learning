package main

import (
	calculate "26_unittest/calculate"
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d is event? %v\n", i, calculate.Even(i))
	}
}
