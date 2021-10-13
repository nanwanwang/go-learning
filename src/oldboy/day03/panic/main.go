package main

import "fmt"


func f1(){
	fmt.Println("f1")
}

func f2(){
	panic("error")
}



func main(){

	defer func(){
	
		if err:=recover();err!=nil{
			fmt.Println("catch panic error")
		}
		fmt.Println("释放数据库连接")
	}()

	 
	f2()
	f1()



}