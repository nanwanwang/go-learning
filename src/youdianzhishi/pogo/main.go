package main

import (
	"flag"
	"pogo/common/logs"
	"pogo/parser"
)

var log = logs.Log

func init() {
	flag.Parse()
}
func main() {
	//if flag.NArg() == 0 {
	//	log.Error("no video url parameter!")
	//	return
	//}
	//url := flag.Arg(0)
	//log.Debug(url)

	xg := parser.Xigua{
		Url: "https://www.ixigua.com/6984386884688577032",
	}
	_, err := xg.GetVideoInfo()
	if err != nil {
		log.Error(err.Error())
		return
	}

}
