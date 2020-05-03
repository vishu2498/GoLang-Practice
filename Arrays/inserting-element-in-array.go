//Inserting an element to an array
package main

import "fmt"

func main() {
	var arr[5] int
	arr[0]=32
	arr[1]=87
	arr[2]=21

	lengthOfArray:=4 //length is defined less by 1 from actual length of array to prevent 'array: out of bounds' error
	index:=3 //position where we want to insert the element
	valueToInsert:=56
	for i:=lengthOfArray;i>index;i-- {
		arr[i]=arr[i-1] //shifting and copying the consecutive values to the next position from index
	}
	arr[index]=valueToInsert //inserting the element at the freed-up index
	lengthOfArray++

	//Final Array
	for key,value:=range arr {
		fmt.Println(key,value)
	}
}
