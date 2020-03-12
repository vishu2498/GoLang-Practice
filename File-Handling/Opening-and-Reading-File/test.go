package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
func main() {
	fileread()
}
func fileread() error {
	//the return type of this function is "error" because if it successfully opens the file and reads it, it will show it's contents, else it will give error.
	file,err:=os.Open("file.txt")
	if err!=nil {
		return err
	}

	defer file.Close() //will close the file at the end of this function

	contents, err := ioutil.ReadFile("file.txt")
	if err!=nil {
		return err
	}
	fmt.Println(string(contents)) //conversion to string is necessary

	return nil
}
