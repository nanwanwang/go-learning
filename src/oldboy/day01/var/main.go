package main

import "fmt"

func main() {
	fmt.Printf("var")

	var (
		name string
		age  int
		isOk bool
	)

	name = "hahah"
	age = 17
	isOk = true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOk)

}
