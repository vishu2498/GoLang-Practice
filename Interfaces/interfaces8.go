package main

import "fmt"

func main()  {
	//Type-switch
	//It is used when we want to find out the data-type of a variable in case if its not defined.
	//It is compulsory to have an interface type variable for type-switch
	var var1 interface{} = "hello"
	switch var1.(type) { //('type' checks for data-type)
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("true or false")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("none")
	}

	//Using type-switch for arrays
	//Even if the data-type of array is same in case and actual definition, the case also checks for the size of array.
	//Since arrays of different sizes are entirely different and have different memory location.
	var var2 interface{}=[4]int{}
	switch var2.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case [3]int:
		fmt.Println("array of 3 integers")
	case [4]int:
		fmt.Println("array of 4 integers")
	}
	fmt.Println()
}
