package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"name"`
}

func main() {

	fmt.Println("欢迎光临学生管理系统")
	fmt.Println(
		`
		1.查看学生信息
		2.新增学生
		3.删除学生
		`)
	// var choice int
	// fmt.Scanln(&choice)
	// fmt.Printf("你选择了%d\n", choice)

	p:= Person{
		"demon",
		17,
	}


	if bytes,err:=json.Marshal(p);err==nil{
		fmt.Println(string(bytes))
	}

}
