package main

import "fmt"

func main(){
	type MyInt1 int //类型定义
	type MyInt2 = int //别名
	var i int = 1
	var a MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Printf("%T %T %T",i,a,i2)
}
