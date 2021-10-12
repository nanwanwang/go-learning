package main

import "fmt"

func main() {
	arr := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
}
