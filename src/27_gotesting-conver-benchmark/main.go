package main

import (
	"27_gotesting-conver-benchmark/calculate"
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d is event? %v\n", i, calculate.Even(i))
	}
}
