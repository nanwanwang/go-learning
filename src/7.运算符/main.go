package main

import "fmt"

func main() {
	a := 1
	a--
	var b int = 10
	c := a + b
	fmt.Println(a)
	fmt.Println(c)

	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

	var f uint = 60
	//var d uint = 13
	//e := f & d
	fmt.Println(f >> 1)
}
