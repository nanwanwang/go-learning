package main

import (
	"fmt"
	"strings"
	"unicode"
)


func wordCount(str  string) int{
	s:=strings.Split(str," ")
    return len(s)
}

func main() {

	str := "hello到底"
	sum := 0
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			sum++
		}
	}
	fmt.Println(sum)


	fmt.Println("-----------")


	fmt.Println(wordCount("how do you do"))
}
