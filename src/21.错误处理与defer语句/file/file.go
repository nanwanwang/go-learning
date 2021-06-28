package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	defer fmt.Println("first defer function")
	defer file.Close()
	defer fmt.Println("second defer function")
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
