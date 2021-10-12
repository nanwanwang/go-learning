package main

import "fmt"

func main() {
	n := 19
	p:= &n
	fmt.Println(p)
	fmt.Printf("%T\n",p)

	var a = new(int)
	*a = 12
	fmt.Printf("%T\n",a)
	fmt.Println(*a)
}