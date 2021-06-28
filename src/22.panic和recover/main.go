package main

import "fmt"

func tryPanic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("catch panic in recover function:%s\n", e)
		}
	}()
	fmt.Println("first line in tryPanic function")
	//panic("call a panic")
	a := 10
	b := 0
	c := a / b
	fmt.Println(c)
	fmt.Println("second line in tryPanic function")
}

func tryPanic2() {
	fmt.Println("first line in tryPanic function")
	panic("call a panic")
	fmt.Println("second line in tryPanic function")
}


type panicFunc func()

func protect (g panicFunc) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("catch panic in recover function:%v\n", err)
		}
	}()
	g()
}
func main() {

	fmt.Println("start call tryPanic")
	tryPanic()
	fmt.Println("end call tryPanic")
	protect(tryPanic2)
}
