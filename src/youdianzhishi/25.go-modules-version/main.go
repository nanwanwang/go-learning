package main

import (
	"fmt"
	stringsV3 "github.com/nanwanwang/gomoduletest/v3/stringsx"
	stringsV2 "github.com/nanwanwang/gomoduletest/v2/stringsx"
)

func main() {
	fmt.Println(stringsV3.Hello("demonwang"))
	if greet, err := stringsV2.Hello("demonwang", "ch"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greet)
	}
}
