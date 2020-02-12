package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type eg1 struct {
	value1 string
}
//Using functions of error package
func main() {
	//Usually when the errors are checked and returned while using 'if err!=nil', the error is returned but without any context i.e. without any additional information except the error name.

	//Default value and data-type of an error is '<nil>'
	var value1 error
	fmt.Printf("%v,%T",value1,value1)
	//'errors.New()' takes an argument and returns it as an error
	var1:=errors.New("testing-error")
	fmt.Println(var1)
	fmt.Println(errors.New("error-printing")) //can be used with 'fmt.Println()' too.
	fmt.Printf("error is: %s or %v and its data-type is: %T",var1,var1,var1) //any of the '%s' & '%v' can be used for printing errors
	fmt.Printf("error details are: %+v",var1) //'%+v' displays error details from 'stackTrace' interface
	fmt.Println()

	var2 := errors.New("1st")
	//The 'errors.Wrap()' function returns a new error that adds context to the original error by recording a stack trace at the point Wrap is called, together with the supplied message.
	//'errors.Wrap()' takes exactly two arguments. The 1st is the previous error message and second will be our custom message in form of a string.
	//It constructs a stack of errors.
	err2 := errors.Wrap(var2, "2nd") //2nd argument will always be placed before the 1st argument in the output.
	fmt.Println(err2)

	var3 := errors.New("whoops")
	err3 := errors.Unwrap(var3) //Unwrap returns the next error in err's chain. If there is no next error, Unwrap returns nil.
	fmt.Println(err3)
	fmt.Println()

	_,err4:=os.Open("file.txt")
	//'errors.Cause()' returns the reason why the error has occurred
	fmt.Println(errors.Cause(err4))
	fmt.Println()

	var4:=errors.New("4th")
	//'errors.Errorf()' performs the similar function of representing errors like 'fmt.Printf()'
	fmt.Println(errors.Errorf("%+v",var4))
	fmt.Println()

	var5:=errors.New("5th")
	fmt.Println(errors.Is(var5,value1))

	//New, Errorf, Wrap, and Wrapf record a stack trace at the point they are invoked.
}