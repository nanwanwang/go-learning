package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func timeDemo() {
	now := time.Now()
	fmt.Println(now)
	//格式化时间用这种方式,后面的年月日要记住
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//时间比较
	prev := time.Date(2020, 10, 10, 17, 20, 34, 12340, time.UTC)
	fmt.Println(prev.Before(now))
	fmt.Println(prev.After(now))
	fmt.Println(prev.Equal(now))
	fmt.Println(now.Sub(prev))
}

func randDemo() {
	fmt.Println(rand.Intn(100))
}

//给随机数添加一个种子
func init() {
	rand.Seed(time.Now().Unix())
}

type BaseResponse struct {
	Code int          `json:"code"`
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonDemo() {
	br := BaseResponse{
		Code: 1,
		Data: ResponseData{
			Name: "demon",
			Age:  20,
		},
	}
	jsonBytes, err := json.Marshal(&br)
	if err == nil {
		fmt.Println(string(jsonBytes))
	}
	var br2 BaseResponse
	_ = json.Unmarshal(jsonBytes, &br2)
	fmt.Println(br2)
}

func regexDemo() {
	input := "My email is nanwanwang@hotmail.com  xxx@abc.com"
	//使用这种方式回返回错误
	exp, err := regexp.Compile("([a-zA-Z0-9]+)@[a-zA-Z0-9]+.[a-zA-Z0-9]+")
	if err == nil {
		fmt.Println(exp.FindString(input))
		fmt.Println(exp.FindAllString(input, -1))
		for _, subMatch := range exp.FindAllStringSubmatch(input, -1) {
			fmt.Println(subMatch[1])
		}
	}
	//不返回错误 如果出错直接panic
	exp2 := regexp.MustCompile("([a-zA-Z0-9]+)@[a-zA-Z0-9]+.[a-zA-Z0-9]+")
	fmt.Println(exp2.FindAllString(input, -1))

}



func main() {
	timeDemo()
	fmt.Println("=============")
	randDemo()
	fmt.Println("=============")
	jsonDemo()
	fmt.Println("=============")
	regexDemo()
}
