package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Day())
	fmt.Println(now.Minute())

	fmt.Println(now.Format("2006-01-02 15:04:05.000"))

	// time.Sleep(1*time.Second)
	// time.Sleep(time.Duration(1))

	// n:=5
	// time.Sleep(time.Duration(n)*time.Second)
	time1, err := time.Parse("2006-01-02 15:04:05", "2021-10-14 15:42:28")
	if err == nil {
		fmt.Println(time1)
	}

	if loc, err := time.LoadLocation("Asia/Shanghai"); err == nil {
		if time2, err2 := time.ParseInLocation("2006-01-02 15:04:05", "2021-10-15 15:42:28", loc); err2 == nil {
			fmt.Println(time2)

			res := time2.Sub(time.Now())
			fmt.Println(res)
		}
	}
}
