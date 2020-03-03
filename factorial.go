package main

import "fmt"

func main() {
	value := 5
	fmt.Println(fact(value))
}
func fact(value int) int {
	f := 1
	for i := 1; i <= value; i++ {
		f = f * i
	}
	return f
}
