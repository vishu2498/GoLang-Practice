package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	cause := errors.New("whoops")
	err := errors.Unwrap(cause) //Unwrap returns the next error in err's chain. If there is no next error, Unwrap returns nil.
	fmt.Println(err)
}