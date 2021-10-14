package main

import "fmt"

type cat struct{}

type dog struct{}

func (c cat) Speak(str string) {
	fmt.Println("喵喵...")
}


func (d dog) Speak(str string){
	fmt.Println("汪汪..")
}

type Speaker interface{
	Speak(str string)
}


func Speak(speak Speaker){
	speak.Speak("")
}


type Any interface {

}


func main(){


	var c1 cat
	var d1 dog
	Speak(c1)
	Speak(d1)
   
	getAnyType("aaa")
	getAnyType(d1)
	getAnyType(&d1)
	getAnyType(12)


}


func getAnyType(any interface{}){
	 switch any.(type){
	 case string:
		fmt.Println("string")
	 case int:
		fmt.Println("int")
	 case dog:
		fmt.Println("dog")
	 case *dog:
		fmt.Println("*dog")
	 default:
		fmt.Println("未知类型")
	 }

}