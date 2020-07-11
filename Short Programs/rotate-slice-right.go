//Rotating a slice to the right (means shifting back elements to the front of slice)
package main

import "fmt"

func main() {
	egslice:=[]int{1,2,3,4,5,6,7}
	number:=3 //rotating with respect to this number
	fmt.Println(f2(egslice,number))
}

func f2(values []int, number int) []int{
	var finalslice []int
	var frontslice []int
	var backslice []int
	frontslice=values[len(values)-number:]
	backslice=values[:len(values)-number]
	finalslice=append(frontslice,backslice...) //'...' literal is used to add a slice to another slice
	values=finalslice
	return values
}
