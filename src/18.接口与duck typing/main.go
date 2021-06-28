package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Speaker interface {
	Say(string)
}

func SpeakAlphabet(speaker Speaker) {
	speaker.Say(";;;;;;;;;;;;;;;;;;")
}

type Person struct {
	name string
}

// Say Person实现了Speaker接口
func (person *Person) Say(message string) {
	fmt.Println(person.name, ":", message)
}

type SpeakWriter struct {
	w io.Writer
}

// Say 实现了Speaker接口
func (sw *SpeakWriter) Say(message string) {
	io.WriteString(sw.w, message)
}

type FileWriter struct {
	filename string
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	err = ioutil.WriteFile(fw.filename, p, 0644)
	return 0, err
}

func main() {
	person := new(Person)
	person2 := Person{
		name: "demon2",
	}
	person.name = "demon"
	person.Say("hello world")

	SpeakAlphabet(person)
	SpeakAlphabet(&person2)

	console := new(SpeakWriter)
	console.w = os.Stdout
	SpeakAlphabet(console)

	fileWriter := new(FileWriter)
	fileWriter.filename = "1.txt"
	fileSpeakWriter := new(SpeakWriter)
	fileSpeakWriter.w = fileWriter
	SpeakAlphabet(fileSpeakWriter)

}
