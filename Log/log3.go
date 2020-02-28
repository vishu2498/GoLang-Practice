package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)
var (
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical problem
)
func init() {
	file, err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
func main() {
	//'log.New()' creates a new custom logger that takes 3 arguments.
	//1st argument is the destination we want the logger to write to. This is provided as a value that implements the 'io.Writer' interface.
	//2nd argument is the prefix for the log similar to 'log.SetPrefix()'.
	//3rd argument is the flag value of log i.e. constants of log package that have the values set as 'syscall' constants having some operational values.
	fmt.Println(log.New(os.Stderr,"ERROR OPEN:",os.O_CREATE))
	fmt.Println()

	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}

