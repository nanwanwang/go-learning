package main

import (
	"fmt"
	"unsafe"
)


type slice struct{
	array unsafe.Pointer
	len int
	cap int
}

func changeELe(s []int){
	s[0] =1000
}
func main() {

	var arr [10]int
	arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:6]
	s2 := arr[:6]
	s3 := arr[2:]
	s4 := arr[:]

	s5 := make([]int, 5, 10)
	fmt.Printf("s5=%v,size=%d,capacity=%d\n", s5, len(s5), cap(s5))

	s6 := arr[2:6]
	fmt.Printf("s6=%v,size=%d,capacity=%d\n", s6, len(s6), cap(s6))
	s7:=s6[3:8]
	fmt.Printf("s7=%v,size=%d,capacity=%d\n", s7, len(s7), cap(s7))
	fmt.Println(s4)
	fmt.Println(s3)
	fmt.Println(s2)
	fmt.Println(s)

	changeELe(s6)
	fmt.Println(s6)
	fmt.Println(arr)
}
