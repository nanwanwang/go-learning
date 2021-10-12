package main

import (
	"fmt"

)

func main(){
	var s1 []int
	var s2 [] string
	fmt.Println(s1,s2)
	fmt.Println(s1==nil)
	fmt.Println(s2==nil)
	s1 = [] int {1,2,3}
	s2 = [] string {"tom","demon","jack"}
	fmt.Println(s1,s2)
   fmt.Println(s1==nil)
   fmt.Println(s2==nil)
   fmt.Println(len(s1),cap(s1))
    

   a1:=[...]int{1,3,5,7,9,11,13}
   s3:=a1[0:4]
   s4:=a1[:]
   s5:=a1[4:]
   s6:=append(s5,3)
   fmt.Println(s3,len(s3),cap(s3))
   fmt.Println(s4,len(s4),cap(s4))
   fmt.Println(s5,len(s5),cap(s5))
   fmt.Println(s6,len(s6),cap(s6))
   fmt.Println(s5,len(s5),cap(s5)) 
   fmt.Println(s4,len(s4),cap(s4))

   s7:= make([]int,0)
   s8:=make([]int,0)
   fmt.Println(len(s7),s7==nil)
   fmt.Println(len(s8),s8==nil)

   arr5:=[...]int{1,2,3,4}
   s9:=arr5[:]
   s8 = append(s8, s9...)
   fmt.Println(s8)
   copy(s7,s9)
   fmt.Println(s7)


   fmt.Println("----")
   s9=append(s9[:1],s9[2:]...)
   fmt.Println(s9)
   fmt.Println(arr5)
}