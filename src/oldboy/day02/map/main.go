package main

import "fmt"

func main() {

	var m map[string]int = make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	
    
	if v,ok:=m["c"];ok{
		fmt.Println(v)
	}else{
		fmt.Println("not found")
	}
    

	delete(m,"a")
	for v, k := range m {
		fmt.Println(v,k)
	}

    
}