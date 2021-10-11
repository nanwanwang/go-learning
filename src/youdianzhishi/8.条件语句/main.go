package main

import "fmt"

func main() {
	a := 5
	switch a {
	case 5:
		fmt.Println("a == 5")
		fallthrough
	case 6,7,8:
		fmt.Println("a == 6")
	case 30:
		fmt.Println("a == 30")
	default:
		fmt.Println("a == default")
	}

	switch {
	case a < 5:
		fmt.Println("a < 5")
	case a == 5:
		fmt.Println("a == 5")
		fallthrough
	default:
		fmt.Println("a > 5")
	}
}
