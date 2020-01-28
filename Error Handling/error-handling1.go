package main

import "fmt"
import "errors"

// divide two number
func Divide(a int, b int) (int, error) {

	// can not divide by `0`
	if b == 0 {
		return 0, errors.New("Can not devide by Zero!")
	} else {
		return (a / b), nil
	}
}

func main() {

	// divide 4 by 0
	if result, err := Divide(6, 0); err != nil {
		fmt.Println("Error occured: ", err)
	} else {
		fmt.Println("6/0 is", result)
	}
}
