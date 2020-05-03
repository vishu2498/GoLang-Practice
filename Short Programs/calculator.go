package main

import "fmt"

func main() {
	var value1, value2, result float64
	var operator string
	fmt.Print("Enter 1st value: ")
	fmt.Scanln(&value1)
	fmt.Println()
	fmt.Print("Enter 2nd value: ")
	fmt.Scanln(&value2)
	fmt.Println()
	fmt.Print("Enter operator: ")
	fmt.Scanln(&operator)
	fmt.Println()
	fmt.Println()

	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		result = value1 / value2
	default:
		panic("no operator entered")
	}

	fmt.Println("So,",value1,operator,value2,":",result)
}
