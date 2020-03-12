package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//'os.Open()' opens files in only read-only mode.

	//'os.OpenFile()' opens files with the option of making certain operations on file (eg. opening, read-only, appending, write-only etc.)
	//If the file does not exist, and the 'O_CREATE' flag is passed to create the file intentionally.
	//It takes arguments as the name of file, flags for operations and file mode.
	file3,err:=os.OpenFile("file3.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file3.Close()
		log.Print("couldn't open the file")
		return
	}

	newLine:="how dare you???"
	_,errwrite:=fmt.Fprintln(file3,newLine)
	if errwrite != nil {
		log.Print("couldn't append to file")
		return
	}
	errclose3 := file3.Close()
	if errclose3 != nil {
		log.Fatal("couldn't close the file")
	}
}
