package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {



	file, err := os.Open("./1.txt")
	defer file.Close()
	if err!=nil{
		fmt.Printf("file open error,err:%v\n",err)
	}
	// if err == nil {
	// 	fmt.Println(file.Name())
	// 	file.WriteString("hahhah")
	// 	var bytes = make([]byte, 100)
	// 	if n, er := file.Read(bytes); er == nil {
	// 		fmt.Println(n)
	// 		fmt.Println(string(bytes))
	// 	}

	// }


	fileobj,err:=os.OpenFile("./1.txt",os.O_TRUNC,0644)
    
	defer fileobj.Close()

	if err ==nil{

		fileobj.Write([]byte("hehehhehehheh\r\n"))
	
	   
	
	}



	if bytes,err:=ioutil.ReadFile("./1.txt");err==nil{
		fmt.Println(string(bytes))
	}




	reader:=bufio.NewReader(file)
	for {

		line,err:=reader.ReadString('\r')
        
  
		if err==io.EOF{
			fmt.Println("读完了..")
			break
		}
		// if len(line) > 0{
		// 	fmt.Println(line)
		// }
		if err!=nil{
			fmt.Printf("read line failed,err:%v",err)
			return
		}else{
			fmt.Println(line)
		}
	}




	

}
