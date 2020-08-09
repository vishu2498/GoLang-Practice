package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"os"
)

func main() {
	variableForWebsite:="google.com"

	//Hitting Google Website
	resp,err:=http.Get("https://www."+variableForWebsite)
	if err!=nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//Reading webpage entire code and storing it inside a variable
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil {
		fmt.Println(err)
	}
	storethis:=string(body)

	//Creating a file for storing the code
	filecreation,err:=os.Create("google.html")
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(filecreation.Name())

	//Writing the code to file
	writingtofile,err:=filecreation.WriteString(storethis)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(writingtofile)
}
