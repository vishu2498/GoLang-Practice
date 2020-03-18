//Two explicit custom error packages are useful for getting more information about errors and performing operation from them.
//1. 'errorx' package
//2. 'go-multierror' package

package main

import (
	"errors"
	"fmt"

	"github.com/goware/errorx"
	"github.com/hashicorp/go-multierror"
)

func main() {
	err1 := errorx.New(500, "error from vishu")
	fmt.Printf("%v,%T", err1, err1)
	fmt.Println()
	fmt.Println(err1.ErrorCode())
	fmt.Println(err1.Message)
	//We can even get the error's cause, details and it's stack.

	var1 := errors.New("got this error")
	err3 := multierror.Append(err1, var1) //Appends the errors like 'errors.Wrap()' but this function takes errors instead of strings.
	fmt.Printf("%v,%T", err3, err3)
	fmt.Println()
	err4 := multierror.Prefix(var1, "what we got:") //adds a custom string to the previous error provided in 1st argument
	fmt.Printf("%v,%T", err4, err4)
	fmt.Println()
	fmt.Println(err3.ErrorOrNil()) //checking if error is nil
	fmt.Println("length of errors: ", err3.Len())

	var var2 error
	var2 = nil
	var var3 error
	var3 = nil
	err5 := multierror.Append(var2, var3)
	fmt.Println(err5.ErrorOrNil())
}