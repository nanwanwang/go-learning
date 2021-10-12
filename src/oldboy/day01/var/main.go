package main

import (
	"fmt"
	"strings"
)

const PI = 3.14

const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota
	a2 = 100
	a3 = iota
	a4 
)

var  s int = 1

func main() {
	fmt.Printf("var")

	var (
		name string
		age  int
		isOk bool
	)

	name = "自负v喊hahah"
	age = 17
	isOk = true
    fmt.Printf("%T\n",age)
	fmt.Printf("%v\n",age)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOk)

	fmt.Println(PI)

    c := 1 + 3i
	fmt.Println(c)
	fmt.Printf("n1=%d,n2=%d,n3=%d\n",n1,n2,n3)
	fmt.Printf("a1=%d,a2=%d,a3=%d,a4=%d\n",a1,a2,a3,a4)

	s:= `C:\Windows\System32\cmd.exe`
	s1:= "hahha"
	ret:=strings.Split(s,"\\")
	fmt.Println(ret)

	s2:=fmt.Sprintf("%s%s",s,s1)
	fmt.Println(s2)



}
