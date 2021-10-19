package main

import (
	"fmt"
	"net"
)

func main() {
	client,err := net.Dial("tcp","localhost:5000")
	if err !=nil{
		fmt.Println("dial error",err)
	}


	var arr =[1024]byte{}
	
	n,err:=client.Read(arr[:])
	if err!=nil{
		fmt.Println("read error: ",err)
	}else{
		fmt.Print(string(arr[:n]) )
	}
	

	


}