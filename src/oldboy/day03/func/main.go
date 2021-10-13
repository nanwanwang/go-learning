package main

import "fmt"

func add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {

	a := 10
	defer fmt.Println("defer")
	if a > 10 {
		fmt.Println("more than 10")

	} else {
		fmt.Println("<= 10")
	}

	fmt.Println(add(2, 3, 4, 5))

}
