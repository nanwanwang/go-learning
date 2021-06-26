package main

import "fmt"

var a1,b1 = 2 ,"hello"
var(
	c1 = 7
	d1 = "world"
)
func main() {
	var a int
	var b string
	c := 5
	var d, e int = 5, 10
	var f, g, h = 1, "string", true

	fmt.Println(a, b, c, d, e, f, g, h)
	fmt.Println(a1,b1,c1,d1)
}
