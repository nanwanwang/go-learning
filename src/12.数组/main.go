package main

import "fmt"

func changeEle(arr [3]int) {
	arr[0] = 1000
}

func changeEleByReference(arr *[3]int) {
	(*arr)[0] = 1000
}

func main() {
	var arr [5]int
	arr = [5]int{1, 2, 3, 4, 5}
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4}
	arr4 := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	fmt.Println(arr)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	fmt.Println(arr3)

	for i, num := range arr3 {
		fmt.Println(i, num)
	}

	for _, num := range arr3 {
		fmt.Println(num)
	}

	for i := range arr3 {
		fmt.Println(i)
	}

	for i := 0; i < len(arr4); i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(arr4[i][j])
		}
	}
	arr5 := [3]int{1, 2, 3}
	changeEle(arr5)
	fmt.Println(arr5)
	changeEleByReference(&arr5)
	fmt.Println(arr5)
}
