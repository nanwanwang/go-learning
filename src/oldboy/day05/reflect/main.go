package main

import (
	"fmt"
	"reflect"
)

func main() {
    i:=1
	typ:=reflect.TypeOf(i)
	fmt.Printf("T=%v\n",typ)

    value:=reflect.ValueOf(i)
	fmt.Println(value.Kind())
	fmt.Println(value)

	


}