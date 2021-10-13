package main

import "fmt"

func f1() {
	fmt.Println("hahh")
}

func f2() int {
	return 1
}

type greeting func(word string) string

func say(greeting greeting, word string) {
	fmt.Println(greeting(word))
}

func english(word string) string {
	return word
}

func main() {
	f := f1
	f3 := f2
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", f3)

	say(english, "hell0")

	f5 := func() {
		fmt.Println("f5")
	}
	f5()
	func() {
		fmt.Println("hell")
	}()
}
