package book

import (
	"fmt"
	"reflect"
)

// book 定义结构体 使用小写表示私有 不可导出 强制使用工厂函数创建
type book struct {
	id      int    "this is a book id"
	title   string "this is a book title"
	author  string "this is a book author"
	subject string "this is a book subject"
}

type techBook struct{
	Cat string
	Title string
	int  //匿名字段
	book  //匿名字段
}

// NewBook 定义构造方法
func NewBook(id int, title, author, subject string) *book {
	return &book{id, title, author, subject}
}

//定义结构体方法
func (book *book) String() string {
	return fmt.Sprintf("id=%d,title=%s,autohr=%s,subject=%s\n\n", book.id,
		book.title, book.author, book.subject)
}

// RefTag 利用反射获取struct 中的Tag
func RefTag(b book, idx int) {
	bType := reflect.TypeOf(b)
	ixField := bType.Field(idx)
	fmt.Printf("%v\n", ixField.Tag)
}

func InitTechBook(){
	bk:=NewBook(1000,"go","demon","about go")
	tb:=new(techBook)
	tb.Cat="tech"
	tb.book.title = "techBook Title"
	tb.int=10
	tb.book = *bk

	fmt.Println("techBook Cat=",tb.Cat)
	fmt.Println("techBook int=",tb.int)
	fmt.Println("techBook book=",tb.book.String())
	//继承book中的String()方法
	fmt.Println(tb.String())
}
