package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用 goroutine和channel 实现一个计算int64随机数各位数和的程序
1.开启一个goroutine循环生成int64类型的随机数,发送到jobChan
2.开启24个goroutine从jobchan中取出随机数计算各位数的和,将结果发送到resultChan
3.主goroutine从resultChan取出结果并打印到终端输出
*/

var wg sync.WaitGroup

type job struct {
	value int64
}

type result struct {
	job    *job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

func generate(c chan<- *job) {

	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		c <- newJob
		time.Sleep(time.Microsecond * 500)
	}

}

func calculate(j <-chan *job, r chan<- *result) {

	defer wg.Done()
	for {
		job := <-j
		sum := int64(0)
		num := job.value
		for num > 0 {
			sum += num % 10
			num /= 10
		}

		newResult := &result{
			job:    job,
			result: sum,
		}

		r <- newResult

	}
}

func main() {

	wg.Add(1)
	go generate(jobChan)

	for i := 0; i < 24; i++ {
		wg.Add(1)
		go calculate(jobChan, resultChan)
	}
	

	for r := range resultChan {
		fmt.Println(r.job.value,r.result)
	}

	wg.Wait()

}
