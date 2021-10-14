package main

import (
	"fmt"
	"oldboy/day07/parser/common"
)

func main() {

	var config common.Config

	if err := common.LoadIni("./conf.ini", &config); err != nil {
		fmt.Println(err)
	}

}
