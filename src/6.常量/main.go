package main

import (
	"fmt"
	"math"
)

const a, b = 3, 4
const (
	c = 4
	d = 5
	MAX = 10
)

const(
	Monday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	const s string = "hello"
	fmt.Println(s)
	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	f := math.Sqrt(a*a + b*b)
	fmt.Printf("%T %T %f", a, f, f)
}
