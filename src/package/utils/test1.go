package utils

import "fmt"

func MyTest2() {
	Count()
}

func init() {

	fmt.Println("test1 init")

    var p = struct{
		name string
		age int
	}{"demon",20}
	

	p1:=struct{
		name string
		age int
	}{"dd",22}
	fmt.Println(p)
	fmt.Println(p1)
}