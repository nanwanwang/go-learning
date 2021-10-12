package main

import "fmt"

func main() {
    s:= "hello 张三"

	for v, k := range s {
		fmt.Printf("%d %c\n",v,k)
	}
}