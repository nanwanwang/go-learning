package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func sendHttpRequestDemo() {
	//构建httpclient 可以设置request 参数
	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if err != nil {
		panic(err)
	}
	//设置User-Agent
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Mobile Safari/537.36 Edg/91.0.864.64")

	http.DefaultClient.Do(request)

	// 基本的http 访问
	//resp, err := http.Get("http://www.baidu.com")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	//sendHttpRequestDemo()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello Golang\n")
	})

	//处理json请求
	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		user := User{
			Name: "demon",
			Age:  20,
		}
		if bytes, err := json.Marshal(user); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		} else {
			writer.Header().Set("Content-type", "application/json")
			writer.Write(bytes)
		}
	})

	//处理图片请求
	http.HandleFunc("/image", func(writer http.ResponseWriter, request *http.Request) {
		image := path.Join("images", "golang.png")
		http.ServeFile(writer, request, image)
	})

	//处理静态html文件
	http.HandleFunc("/html", func(writer http.ResponseWriter, request *http.Request) {
		user := User{
			Name: "demon",
			Age:  20,
		}
		htmlFile := path.Join("templates", "index.html")
		tmpl, err := template.ParseFiles(htmlFile)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(writer, user); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	log.Fatal(http.ListenAndServe(":12345", nil))
}
