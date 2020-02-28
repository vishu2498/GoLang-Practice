/*To separate logging and output in the final display, UNIX architects added a device called 'stderr'.
This device was created to be the default destination for logging. It allows developers to separate
their programs’ output from their logging. For a user to see both the output and the
logging when running a program, terminal consoles are configured to display what’s
written to both 'stdout' and 'stderr'.*/

package main

import (
	"log"
)
//'Log' package is used for specific logging purposes that differentiates it from the output of the program from 'fmt'.
//The purpose of logging is to get a trace of what the program is doing, where it’s happening, and when.

func init() {
	log.SetPrefix("VISHU LOG:") //This function sets the given string to the starting of every log given by the program.
	//This is normally written in capital letters.
}

func main()  {
	log.Print("1st log") //prints the given log string
	log.Println("2nd log") //prints the given log string with new line
	log.Printf("%v, %T","log format","log format") //formatted printing of log
	log.Fatal("Fatal log") //similar to 'log.Print()' followed by a call to os.Exit(1) (exit code is hard-coded)
	//So, anything written after 'log.Fatal()' will not execute as the program will stop execution with exit code 1.
	//Also, log.Fatalf() and 'log.Fatalln()' can be used.
}

