package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" bson:"b_name"`
	Age  int    `json:"age"  bson:"b_age"`
}

func (u User) String() {
	fmt.Printf("Name=%s,Age=%d\n", u.Name, u.Age)
}

// Print 使用指针的话 使用反射不会获取到改方法
func (u User) Print(prex string) {
	fmt.Println(prex + "hello world!")
}

// 基础用法
func reflectBasicUse() {
	u := User{"demon", 22}
	//获取类型
	t := reflect.TypeOf(u)
	fmt.Println(t)

	//获取值
	v := reflect.ValueOf(u) //返回的是一个值的拷贝
	fmt.Println(v)
	fmt.Println(v.Type()) //等效与reflect.TypeOf()

	//获取底层的数据类型
	fmt.Println(v.Type().Kind())
}

// 反射获取方法
func reflectMethod() {
	u := User{"demon", 22}
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("fieldindex:%d,fieldname:%s\n", f.Index, f.Name)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("methodindex:%d,methodname:%s\n", m.Index, m.Name)
	}
}

// 利用反射修改值
func reflectChangeFieldValue() {
	x := 2
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(100)
	fmt.Println(x)
}

// 利用反射获取tag
func reflectStructTag() {
	var u User
	h := `{"name":"demon","age":20}`
	err := json.Unmarshal([]byte(h), &u)
	if err == nil {
		fmt.Println(u)
	} else {
		fmt.Println(err)
	}

	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Tag,f.Tag.Get("json"),f.Tag.Get("bson"))
	}
}

// 利用反射动态调用方法
func reflectDynamicCallMethod() {
	u := User{"demon", 22}
	t := reflect.ValueOf(u) //使用valueof 而不是typeof
	printM := t.MethodByName("Print")

	//判断函数是否存在
	if printM.IsValid() {
		args := []reflect.Value{reflect.ValueOf("PrintPrefix")}
		//args:= []reflect.Value{}
		fmt.Println(printM.Call(args))
	}
}
func main() {
	reflectBasicUse()
	reflectMethod()
	reflectChangeFieldValue()
	reflectDynamicCallMethod()
	reflectStructTag()
}
