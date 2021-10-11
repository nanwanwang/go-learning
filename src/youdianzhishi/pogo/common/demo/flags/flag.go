package flags

import (
	"flag"
	"fmt"
	"os"
)


func main(){
	fmt.Println("I'm Pogo Project")
	flagInt:=flag.Int("flagname",100,"help message for flagname")
	flagString:=flag.String("flagstr","abc","help message for flag string")

	var flagValue int
	flag.IntVar(&flagValue,"flagname2",132,"help message for flagname2")
	flag.Usage= func() {
		fmt.Fprintf(os.Stderr,`pogo version: pogo/1.0.0
Usage: pogo [-h] [-p prefix]
Options:
`)
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println(flag.NFlag())
	fmt.Println(*flagInt,*flagString,flagValue)

	//flag.Value()

	//接收非flag参数值
	fmt.Println(flag.Arg(0))
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())

	// 传参数的两种方式 -flag= 值 或者 -flag 空格  值
	//go run main.go -flagname=300 -flagstr adddd

}


