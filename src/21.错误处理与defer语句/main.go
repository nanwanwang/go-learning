package main

import (
	"21.错误处理与defer语句/file"
	"errors"
	"fmt"
)

type Divide struct {
	dividee int
	divider int
}

func (d *Divide) Error() string {
	strFormat := `
     Cannot proceed,the divider is zero.  
     dividee:%d
     divider:0`
	return fmt.Sprintf(strFormat, d.dividee)
}

func ComputeDiv(d *Divide) (result int, err error) {
	if d.divider == 0 {
		err = d
	} else {
		result = d.dividee / d.divider
	}
	return result, err
}

func main() {
	err := errors.New("a new err object")
	fmt.Printf("%v\n", err)

	err = fmt.Errorf("a fmt error format object:%s", err.Error())
	fmt.Printf("%v\n", err)

	d := new(Divide)
	d.dividee = 10
	d.divider = 2
	if result, err := ComputeDiv(d); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}

	if content, err := file.ReadFile("1.txt"); err == nil {
		fmt.Println(content)
	}else{
		fmt.Println(err.Error())
	}

}
