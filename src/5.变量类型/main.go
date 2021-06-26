package main

import (
	"fmt"
	"runtime"
	"strconv"
	"unsafe"
)

func main(){
	var a = 64
	var b int32 = 32
	var c string = "hello world"
	var d byte = 128
	var e bool = false
	var f rune = '中' //跟int32一样
	var g float64 = 45.6
	var h complex128 = 3
	var i int32
	i = int32(a)
	ptr := &a
	fmt.Printf("%T\n",ptr)
	cpuArch:= runtime.GOARCH
	intSize:=strconv.IntSize
	fmt.Println(cpuArch,intSize,i)
	fmt.Printf("%T %T %T %T %T %T %T %T\n",a,b,c,d,e,f,g,h)
	fmt.Println(g)
	fmt.Println(unsafe.Sizeof(a),unsafe.Sizeof(b),unsafe.Sizeof(c),
		unsafe.Sizeof(d),unsafe.Sizeof(e),unsafe.Sizeof(f),unsafe.Sizeof(g),unsafe.Sizeof(h))
}
