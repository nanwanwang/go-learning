package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp","127.0.0.1:5000")
	if err !=nil{
		fmt.Println("start tcp server on 127.0.0.1:5000 failed,err:",err)
	}

	fmt.Println("server start...")

	conn,err:=listener.Accept()
	if err!=nil{
		fmt.Println("accept failed,err:",err)
	}

	n,err:=conn.Write([]byte("hello"))
	if err !=nil{
		fmt.Println("write error")
	}

	fmt.Println(n)


	
}