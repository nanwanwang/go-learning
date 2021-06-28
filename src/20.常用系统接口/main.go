package main

import (
	"20.常用系统接口/example"
	"fmt"
)

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

type List []int

func (list *List) Append(element int) {
	*list = append(*list, element)
}

func (list List) Len() int {
	return len(list)
}

func Counter(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func IsLarge(l Lener) bool {
	return l.Len() > 50
}

func main() {
	var list List
	Counter(&list, 1, 10) //接收者是引用类型,不可以传入值类型
	IsLarge(list)

	//指针
	plist := new(List)
	Counter(plist, 1, 10) //接收值类型时,传入可以是指针,会解引用
	IsLarge(plist)

	course := new(example.Course)
	course.Title = "go"
	course.SubTitle = "about go"
	fmt.Println(course)
	fmt.Printf("%v\n", course)

	intArray := example.IntArray{1, 3, 4, 5, 45, 6, 99, 34}
	example.Sort(&intArray)
	fmt.Println(intArray)
}
