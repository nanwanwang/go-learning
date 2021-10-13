package main

import (
	"fmt"
)


func modify(arr *[5]int){
	arr[1] = 100
}

func main() {
	arr := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range arr {
		sum += v
	}


	arr2:=arr
	arr2[0]= 100
	fmt.Println(arr2)
	fmt.Println(arr)
	fmt.Println(sum)

	var arr3 [5]int
	
	fmt.Println(arr3)

	fmt.Println("----------------")
	modify(&arr)
	fmt.Println(arr)


	var arr4 [4]int
	var arr5 [4] int
	fmt.Println(arr4==arr5)
	fmt.Println(len(arr4))

	var str  string
	fmt.Println(str=="")


}
