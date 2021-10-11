package main

import "fmt"

type Phone interface {
	Call(string)
}

type Camera interface {
	Take() string
}

// SmartPhone 接口嵌套
type SmartPhone interface {
	Phone
	Camera
	Download(string) string
}

type MiPhone struct {
	Logo string
}

func (mp *MiPhone) Call(phone string) {
	fmt.Println("Call to phone:", phone)
}

func (mp *MiPhone) Take() string {
	return "logo.png"
}

func (mp *MiPhone) Download(url string) string {
	return fmt.Sprintf("visit the url:%s", url)
}

func ListSmartPhoneFunction(sp SmartPhone) {
	//类型判断
	if v, ok := sp.(*MiPhone); ok {
		v.Call("10086")
		fmt.Println("sp.Take()", v.Take())
		fmt.Println("sp.Download()", v.Download("https://wwww.baidu.com"))
	} else {
		fmt.Println("not MiPhone Pointer type")
	}

}

// Any 空接口 所有接口都实现空接口 类似C#中 object基类  等同于 interface{}
type Any interface {
}

// GetAnyType 空接口作为参数可以接收任何数据类型
func GetAnyType(any interface{}) {
	switch t := any.(type) {
	case int:
		fmt.Println("any is int type")
	case string:
		fmt.Println("any is string type")
	case *MiPhone:
		fmt.Println("any is MiPhone pointer type")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}

// Queue  定义队列
type Queue []interface{}

func (queue *Queue) Push(element ...interface{}) {
	*queue = append(*queue, element...)
}

func (queue *Queue) Pop() interface{} {
	head := (*queue)[0]
	*queue = (*queue)[1:]
	return head

}

func main() {
	miPhone := new(MiPhone)
	miPhone.Logo = "xiaomi"
	ListSmartPhoneFunction(miPhone)

	var va1 Any
	va1 = 5
	fmt.Printf("va1 value :%v \n", va1)

	str := "abcd"
	va1 = str
	fmt.Printf("va1 value:%v \n", va1)

	va1 = miPhone
	fmt.Printf("va1 value:%v \n", va1)

	GetAnyType(str)     //any is string type
	GetAnyType(miPhone) //any is MiPhone pointer type
	GetAnyType(2.1)     //Unexpected type float64

	queue := new(Queue)
	queue2 := Queue{1, 2, 3, 5}
	fmt.Println(queue2)
	queue.Push("a", "b", 3, 4, 5)
	fmt.Println(queue)
	queue.Pop()
	fmt.Println(queue)
}
