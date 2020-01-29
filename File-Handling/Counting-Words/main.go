package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//"contents" variable stores the ASCII values of each character
	text := string(contents)           //converting & storing data as a string from the file
	count := len(strings.Fields(text)) //counts words in the file (actually the variable storing the string converted from ASCII)
	fmt.Println(count)
}

