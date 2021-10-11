package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:6]
	fmt.Println(s1, len(s1), cap(s1))
	//append追加的时候如果底层数组容量够则改变底层数组,不够回创建
	//新的数组,原来底层数组的内容不会发生改变
	s1 = append(s1, 10, 20, 30, 40)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(arr)

	//使用copy 向s3中拷贝元素如果s3长度比s2小则复制s3长度的内容
	s2 := []int{1, 2, 3, 4}
	s3 := make([]int, 3)
	fmt.Println(s2, s3)
	copy(s3, s2)
	fmt.Println(s2, s3)

	//使用append删除元素4
	s4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s5 := s4[:4]
	s6 := s4[5:]
	//使用 ... 可以接收slice参数
	s5 = append(s5, s6...)
	fmt.Println(s5)

	//删除第一个元素
	s6 = s6[1:]
	fmt.Println(s6)

	var s10 []int //该初始化为nil
	var s11 = make([]int,0) //该初始化不为nil

	if s10 == nil{
		fmt.Println("s10 == nil")
	}else {
		fmt.Println("s10 != nil")
	}
	s10=append(s10,1,2,3)
	fmt.Println(s10)

	if s11 == nil{
		fmt.Println("s11 == nil")
	}else {
		fmt.Println("s11 != nil")
	}
}
