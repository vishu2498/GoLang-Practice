package main

import (
	"fmt"
)

func main() {
	//Switch case
	//The executing statements inside 'case' do not need curly braces
	switch true {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	default: //used when all other cases fail
		fmt.Println("none")
	}

	//Using multiple cases in single case
	var1:=50
	switch var1 {
	case 51,64,89:
		fmt.Println("not equal")
	case 50,48,32:
		fmt.Println("somewhat equal")
	default:
		fmt.Println("none")
	}
	//Duplicate cases are not allowed. It means that cases which have the same values will throw an error

	//Using initializer syntax in switch-case
	//Variables initialized within initializer syntax of switch-case could not be used elsewhere
	switch var2:=2*3; var2{
	case 2:
		fmt.Println("wrong")
	case 6:
		fmt.Println("right")
	}

	//Taglist syntax
	//In this syntax, the variable on which the cases are checked is not written with switch statement but with the cases
	var3:=45
	switch {
	//So, the checking is done by writing the conditions with cases
	case var3<10:
		fmt.Println("lesser than 10")
	case var3>10:
		fmt.Println("greater than 10")
	}
	//In taglist syntax, in case if two or more cases refer to the same result (eg. value is 5 and 1st case checks <10 and 2nd case checks <20), then the 1st case of all of them will execute.

	//'Fallthrough' statement is used when the case just below the case where 'fallthrough' is written also needs to be executed
	//It doesn't care if the next case passes or fails (logic-less). It will execute the statement in any condition.
	var4:="welcome"
	switch var4 {
	case "welcome":
		fmt.Println("welcome is present")
		fallthrough
	case "hello":
		fmt.Println("welcome is also present")
	}

	//Type-switch
	//It is used when we want to find out the data-type of a variable in case if its not defined.
	//It is compulsory to have an interface type variable for type-switch
	var var5 interface{} = "hello"
	switch var5.(type) { //('type' checks for data-type)
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
	var var6 interface{}=[4]int{}
	switch var6.(type) {
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

	//'break' keyword
	//It is used in a case when there are multiple statements in a case and we want to execute only few of them
	switch true {
	case true:
		fmt.Println("true")
		fmt.Println("really true")
		break //It will stop the execution of case till this line
		fmt.Println("damn true")
	case false:
		fmt.Println("false")
	}
}

