package common

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

func LoadIni(fileName string, data interface{}) (err error) {

	//0.参数的校验
	//0.1 传进来的data参数必须是指针类型(因为需要在函数中对其赋值)
	//0.2 传进来的data参数必须是结构体类型指针(因为配置文件中各种键值对需要赋值给结构体的字段)

	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("data should be a pointer")
		return
	}

	if t.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("data should be a struct pointer")
		return
	}
	//1.读文件
	//2.一行一行读
	//2.1 如果是注释就跳过
	//2.2 如果是中括号[开头就表示节
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	fileContent := string(fileBytes)
	strSlice := strings.Split(fileContent, "\r\n")

	fmt.Println(strSlice)

	var structName string
	for idx, line := range strSlice {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			if !strings.HasPrefix(line, "[") || !strings.HasSuffix(line, "]") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}

		} else {
			//2.3 如果不是中括号[开头就是分割的键值对
			if !strings.Contains(line, "=") || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d synatx error", idx+1)
				return
			}
			
			v:=reflect.ValueOf(data)
			structObj:=v.Elem().FieldByName(structName)
			if structObj.Kind()!=reflect.Struct{
				err= fmt.Errorf("data中的%s字段应该是一个结构体",structName) 
				return
			}

			strings.Split(line, "=")
		}

	}

	return

}
