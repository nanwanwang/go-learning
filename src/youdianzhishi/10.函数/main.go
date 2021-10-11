package main

import (
	"fmt"
	"math"
)

//多返回值
func operator(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("not support operate:%s\n", op)
	}
}

//多返回值命名
func swap(a, b int) (x, y int) {
	x, y = b, a
	return
}

//传递函数
func compute(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func powInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//定义类型
type greeting func(name string) string

func say(greeting greeting, word string) {
	fmt.Println(greeting(word))
}

func english(name string) string {
	return "hello " + name
}

//可变参数
func sum(nums ...int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}

func main() {

	if c, err := operator(1, 3, "x"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(c)
	}

	fmt.Println(swap(2, 2222))

	fmt.Println(powInt(3, 4))

	say(english, "world")

	fmt.Println(sum(1,2,3,4,5,6))
}
