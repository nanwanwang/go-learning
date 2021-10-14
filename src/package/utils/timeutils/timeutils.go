package timeutils

import (
	"fmt"
	"time"
)

func PrintTime() {
	fmt.Println(time.Now())
}

func init() {
	fmt.Println("init()包初始化操作...")
}



func init(){
	fmt.Println("another init")
}