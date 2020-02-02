package main

import "fmt"

func main() {
	//For-range loop is similar to python for-range loop
	//This can be mainly used to iterate over data and collections that don't have their size defined at compile-time
	slice1:=[]int{2,4,8}
	for key,value:=range slice1{ //the 'key' and 'value' variables store the index and data of slice.
	//It is compulsory to write collection name after 'range' keyword.
		fmt.Println(key,value)
	}
	fmt.Println()

	//Using for-range loop for maps
	map1:= map[string]int{
		"1st": 56,
		"2nd": 94,
	}
	for key,value:=range map1{
		fmt.Println(key,value)
	}
	fmt.Println()

	//Using for-range loop for strings
	//It will print out the index and ASCII values of characters of the string
	str1:="vishu"
	for key,value:=range str1{
		fmt.Println(key,value)
	}
	fmt.Println()
	//To get the characters of string directly, we have to type-cast the ASCII value to its correspondent character
	for key,value:=range str1{
		fmt.Println(key,string(value))
	}
	fmt.Println()
	
	//While dealing with maps and for-range loop, we can just directly get keys by just defining the 'key' in for-range loop.
	//However, we can't directly get the values because they are attached to their respective keys
	//For this situation, we can use blank identifier (_)
	map2:=map[string]int{
		"1st": 98,
		"2nd": 27,
	}
	for _,value:=range map2 {
		fmt.Println(value) //getting just values
	}
}
