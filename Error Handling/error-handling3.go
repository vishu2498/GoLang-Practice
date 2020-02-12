package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	cause := errors.New("1st")
	err := errors.Wrap(cause, "2nd")
	fmt.Println(err)
}
