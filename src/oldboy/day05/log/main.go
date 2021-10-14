package main

import (
	"fmt"
	"log"
	"path"
	"runtime"
)

func main() {

	log.Println("测试日志")

	f1()
}

func f1() {
	getInfo()
}

func getInfo() {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("runtime.Caller() invoke failed!")
		return
	}

	//获取调用的函数名称
	funcName := runtime.FuncForPC(pc).Name()

	fmt.Println(funcName)

	//获取函数文件全路径
	fmt.Println(file)
	//获取调用函数所在的文件
	fmt.Println(path.Base(file))
	//获取调用函数躲在的行数
	fmt.Print(line)
}
