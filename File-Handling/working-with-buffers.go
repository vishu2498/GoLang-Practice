package main

import (
	"bytes"
	"fmt"
	"os"
)

func main()  {
	//We will first create a buffer value and write something to it.
	//For this, Write() method of 'io.Writer' interface is used.
	//Write() method takes the data in form of slice of bytes. So, everything needs to be type-casted first to slice of bytes.
	//It returns the number of characters (length) given in its argument and an error.
	//If the buffer becomes too large, Write() will panic with ErrTooLarge.
	var buff bytes.Buffer //this variable is now of type io.Writer
	buff.Write([]byte("vishu"))

	//fmt.Fprintf() is used to concatenate (append) a string to the buffer. It is a variadic function.
	//It takes 3 arguments.
	//1st argument is the address of the value to which the appending is to be done. It accepts values from types that implement the io.Writer interface.
	//2nd argument is what string we want to add.
	//3rd argument is actually an interface so many more string can be added. It is not compulsory.
	fmt.Fprintf(&buff," new")
	//fmt.Fprint() & fmt.Fprintln() are also used for similar purposes.

	//Now, since the new data is added to buffer, we need to finally add it to actual data i.e. stdout device.
	buff.WriteTo(os.Stdout) //stdout is responsible for displaying the output in console
}