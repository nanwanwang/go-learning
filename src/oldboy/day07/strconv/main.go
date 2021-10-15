package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := 65

	s := "1000"


	c,err:=strconv.Atoi(s)
	if err ==nil{
		fmt.Printf("%T,%d\n",c,c)
	}
	

	fmt.Println(s)

	s2:=fmt.Sprintf("%d",num)

	fmt.Println(s2)
	// fmt.Println(string(num))

	i:=97
	// fmt.Println(string(i))
	
}