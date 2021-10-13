package main

import "fmt"

type Book struct {
	id      int
	title   string
	author  string
	subject string
} 


func NewBook(id int,title,author,subject string) Book{
	return Book{
		id,
		title,
		author,
		subject,
	}
}


func editBook(book *Book){
	book.author="hahah"
}

func(book *Book) toString(){
	fmt.Println(book)
}

func main() {

	var book1 Book
	book1.id = 1
	book1.title = "golang"
	book1.subject = "about go"
	book1.author = "demon"

	fmt.Println(book1)


	var book2= new(Book)
	book2.id=2
	book2.title="c#"
	book2.subject="about c#"
	book2.author="demon"
	fmt.Printf("%v\n",book2)
	fmt.Printf("%p\n",book2)

	fmt.Println(&book2)


	book3:=Book{
		id: 3,
		title: "c",
		subject: "about c",
		author: "demon",
	}

	fmt.Println(book3)
	fmt.Printf("%p\n",book3)


	book4:=NewBook(5,"c++","demon","about c++")

	fmt.Println(book4)
	editBook(&book4)
	fmt.Println(book4)


	book4.toString()  
	
}