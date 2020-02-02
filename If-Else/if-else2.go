package main

import "fmt"

func main() {
	//If-else has a concept of short-circuiting.
	/*In short-circuiting, if the if-else conditions finds out that an OR operator is used and if any condition is returning true,
	then its not gonna move on to other conditions for checking. It will simply pass on to the executing statement.*/
	if false || returning() || !true { //Here, when GoLang finds out that the first operation is returning the result that is acceptable, then it will not move on to further OR cases.
		fmt.Println("every condition is false") //this will not be printed since 'if' only accepts true conditions
	}

	//Using 'else'
	var1:=54
	if var1<53 {
		fmt.Println("small")
	} else { //syntax of 'else' is like this only
		fmt.Println("big")
	}

	//Using 'else-if'
	var2:=50
	if var2<49 {
		fmt.Println("small")
	} else if var2<48 {
		fmt.Println("too small")
	} else {
		fmt.Println("big")
	}
}
func returning() bool {
	return false
}
