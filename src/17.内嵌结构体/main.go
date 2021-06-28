package main

import (
	"17.内嵌结构体/book"
	"fmt"
	"strconv"
)

type Integer int

func (it Integer) String() string {
	return strconv.Itoa(int(it))
}

func main() {
	bk := book.NewBook(1000, "go", "demon", "about go")
	fmt.Println(bk.String())
	book.RefTag(*bk, 1)
	book.InitTechBook()
	it:=Integer(5)
	fmt.Println(it.String())
}
