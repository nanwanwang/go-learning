package main

import "fmt"

// book 定义结构体 使用小写表示私有 不可导出 强制使用工厂函数创建
type book struct {
	id      int
	title   string
	author  string
	subject string
}

func printBook(book book) {
	book.id = 1009
}

func printBook2(book *book) {
	book.id = 10010
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

func main() {
	//结构体初始化
	var book1 *book
	//使用new返回指针类型
	book1 = new(book)
	book1.id = 1001
	book1.title = "golang"
	book1.author = "demon"
	book1.subject = "about go"
	fmt.Println(*book1)

	book2 := book{
		id:      1002,
		title:   "python",
		author:  "james",
		subject: "about python",
	}
	fmt.Println(book2)

	//该初始化可以省略字段,但顺序要跟结构体一样
	book3 := book{
		1003,
		"csharp",
		"james",
		"about csharp",
	}
	fmt.Println(book3, book3.id)

	//直接传递结构体是值传递拷贝一份给方法
	printBook(book3)
	fmt.Println(book3)

	//传递指针可以改变值
	printBook2(&book3)
	fmt.Println(book3)

	//调用结构体方法
	fmt.Println(book3.String())

	//使用NewBook构造方法创建book
	book4 := NewBook(1005, "java", "gsl", "about java")
	fmt.Println(*book4)

}
