package main

import (
	"fmt"
	"sort"
)

func main() {
	//初始化 3 种方式
	var m1 map[string]int      //m1 == nil 不能存放值
	m2 := make(map[string]int) //可以存放值
	m1 = make(map[string]int)  //可以存放值
	m1["a"] = 10

	m3 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	if m1 == nil {
		fmt.Println("m1 == nil")
	} else {
		fmt.Println("m1 != nil")
	}
	if m2 == nil {
		fmt.Println("m2 == nil")
	}
	fmt.Println(m3["a"])

	if val, ok := m1["key"]; ok {
		fmt.Println(val)
	}

	//可以使用range来迭代
	for key := range m3 {
		fmt.Printf("key=%s,val=%d\n", key, m3[key])
	}

	for key, value := range m3 {
		fmt.Printf("key=%s,val=%d\n", key, value)
	}

	//使用delete删除
	delete(m3, "b")
	for key, value := range m3 {
		fmt.Printf("key=%s,val=%d\n", key, value)
	}

	m3["b"] = 2
	m3["d"] = 4
	//打印出排序后的map值
	var keys []string
	for key := range m3 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	for _, key := range keys {
		fmt.Println(m3[key])
	}

}
