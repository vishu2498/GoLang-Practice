//Finding how many pairs of 'xy' are there in a string
package main

import "fmt"

func main() {
	value1:="xyxyxyxy"
	value2:="xxyyyxxxxyy"
	value3:="xyyyyyyy"
	fmt.Println(count(value1))
	fmt.Println(count(value2))
	fmt.Println(count(value3))
}

func count(value string) int {
	var counter1 int //Number of 'x' in string
	var counter2 int //Number of 'y' in string
	var finalcounter int //Number of pairs of 'xy' in string

	//Iterating over string to find number of 'x'
	for i:=0;i<len(value);i++ {
		if string(value[i])=="x" {
			counter1++
		}
	}

	//Iterating over string to find number of 'y'
	for i:=0;i<len(value);i++ {
		if string(value[i])=="y" {
			counter2++
		}
	}

	//Number of pairs will be equal to minimum of number of 'x' or 'y' in the string
	if counter1<=counter2 {
		finalcounter=counter1
	} else if counter2<=counter1 {
		finalcounter=counter2
	}

	return finalcounter
}
