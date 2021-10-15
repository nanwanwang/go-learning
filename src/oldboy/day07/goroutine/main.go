package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

func getRandNum() {

	rand.Seed(time.Now().UnixNano())

	r1 := rand.Int()
	r2 := rand.Intn(10)

	fmt.Println(r1, r2)
}

var wg sync.WaitGroup

func main() {

	// for i:=0;i<100;i++{
	// 	a:=i
	// 	go func (){
	// 		fmt.Println(a)
	// 	}()
	// }
	// go hello()

	// getRandNum()

	// fmt.Println("main")

	// //time.Sleep(time.Second)

	// ran:=rand.Intn(5)

	// time.Sleep(time.Millisecond*time.Duration(ran))

	//
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}

	wg.Wait()

	// time.Sleep(time.Second)
}

func f1(i int) {

	defer wg.Done()
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)

}
