package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type wincpauthrequest struct {
	RequestType string
	UserName    string
	DynamicPwd  string
}

func httpGet(request wincpauthrequest) int{
	 response,err:=http.Get("http://localhost:18189/api/wincpauth?RequestType="+request.RequestType+"&UserName="+request.UserName)
	 if err!=nil{
		return  -99
	}
	defer response.Body.Close()
	 body,err:=ioutil.ReadAll(response.Body)
	 if err !=nil{
	 	return  -99
	 }
	 fmt.Println(string(body))
	 return 1
}

func main()  {

	lenArgs := len(os.Args)
	if lenArgs == 0 {
		fmt.Println(11)
	}
	request := wincpauthrequest{
		RequestType: os.Args[1],
		UserName:    os.Args[2],
	}
	if lenArgs == 4 {
		request.DynamicPwd = os.Args[3]
	}
	result:=httpGet(request)

	fmt.Println(result)
}
