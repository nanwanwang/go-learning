package main

import "fmt"

var g int = 10

func sum(a, b int) int {
	var c = 10
	s := a + b + c
	return s
}
func main() {
	var a, b, c int
	a = 10
	b = 20
	c = 30
	var g int = 100
	fmt.Println(a, b, c, g)

	var p int = 20
	var ip *int
	ip = &p
	*ip = 30
	fmt.Println(p)

}
