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
//In majority of the functions, the 1st argument is that error on which we want operations to occur.
func main() {
	//Usually when the errors are checked and returned while using 'if err!=nil', the error is returned but without any context i.e. without any additional information except the error name.

	//Default value and data-type of an error is '<nil>'
	var value1 error
	fmt.Printf("%v,%T",value1,value1)
	fmt.Println()

	//'errors.New()' takes an argument and returns it as an error. It can only take one argument.
	var1:=errors.New("testing-error")
	fmt.Println(var1)
	fmt.Println(errors.New("error-printing")) //can be used with 'fmt.Println()' too.
	fmt.Printf("error is: %s or %v and its data-type is: %T",var1,var1,var1) //any of the '%s' & '%v' can be used for printing errors
	fmt.Println()
	fmt.Printf("error details are: %+v",var1) //'%+v' displays error details from 'stackTrace' interface
	fmt.Println()

	var2 := errors.New("1st")
	//The 'errors.Wrap()' function returns a new error that adds context to the original error by recording a stack trace at the point Wrap is called, together with the supplied message.
	//'errors.Wrap()' takes exactly two arguments. The 1st is the previous error message and second will be our custom message in form of a string.
	//It constructs a stack of errors.
	err2 := errors.Wrap(var2, "2nd") //2nd argument will always be placed before the 1st argument in the output.
	fmt.Println(err2)
	fmt.Println()
	//Whenever "%+v" is used with 'errors.Wrap()' for output, it displays stack traces of both the 1st error and 2nd added message error (string).
	fmt.Printf("%+v",err2)
	fmt.Println()

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

	//'error.Is()' returns a boolean value showing that are two given values in arguments are equal
	var5:=errors.New("5th")
	fmt.Println(errors.Is(var5,value1))
	fmt.Println()

	//'errors.WithMessage()' is similar to 'errors.Wrap()' i.e. it takes the 1st error as argument and adds a custom message in form of string given as 2nd argument.
	//But the difference is that 'errors.WithMessage()' when used with "%+v" will give stack trace of just the 1st error given as argument.
	//Whereas 'errors.Wrap()' used with "%+v" will give stace trace of both the 1st and 2nd errors.
	var6:=errors.New("6th")
	var7:=errors.WithMessage(var6,"7th")
	fmt.Println(var7)
	fmt.Printf("%+v", var7) //stack trace of only var6 error will be displayed
	fmt.Println()

	//'errors.WithMessagef()' is an extension over 'errors.WithMessage()' and similarly adds a custom message with the given error.
	//1st argument is the error dealing with, 2nd is the format in which we will specify the new message, 3rd argument will be the new custom message.
	//Since this function has 3rd argument of type interface{}, it allows to add message of any data-type unlike only string type with 'errors.WithMessage()'.
	var8:=errors.New("8th")
	var9:=errors.WithMessagef(var8,"%s","9th")
	fmt.Println(var9)
	//3rd argument can also be used to take the data-type of the given original error (1st argument)
	var10:=errors.WithMessagef(var8,"%T",var8)
	fmt.Println(var10)

	//'error.As()' checks that the error given in the 1st argument of the function contains the error given by the 2nd argument in its stack trace.
	//If it finds that error in its stack trace, it returns true. Else returns false.
	//This function checks the stack trace by calling 'errors.Unwrap()' multiple times in the error's chain.
	if _, err := os.Open("file.txt"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}



	//New, Errorf, Wrap, and Wrapf record a stack trace at the point they are invoked.
}