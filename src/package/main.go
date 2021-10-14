package main

// import "package/utils"  src/package/utils  绝对路径 导入包  src省略
import (
	"fmt"
	p "package/person"
	"package/pk1"
	"package/utils" //相对路径 相对main.go 文件
    "package/utils/timeutils"
)

func main() {

	utils.Count()
	timeutils.PrintTime()

	pk1.MyTest1()
	utils.MyTest2()

	p1:= p.Person{
		Name: "demon",
		Age: 30,
		Sex: "male",
	}

	fmt.Println(p1)
}
