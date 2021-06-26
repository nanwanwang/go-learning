package main

import (
	"17.内嵌结构体/book"
	"fmt"
)

func main(){
	bk:=book.NewBook(1000,"go","demon","about go")
	fmt.Println(bk.String())
	book.RefTag(*bk,1)
	book.InitTechBook()
}
