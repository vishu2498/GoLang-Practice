package main

import "fmt"

func main() {
	//'&' is address-of operator
	var1:=10
	var2:=&var1 //var2 is holding the memory address of var1
	//var2 is actually 'var var2 *int = &var1'
	fmt.Println(var1,var2)
	//'*' is de-referencing operator
	fmt.Println(*var2) //getting the actual value that is stored in memory location represented by var2
	*var2=25
	fmt.Println(var1) //changing value from var2 changed the value of var1 since they both are pointing to the same data

	//Individual values of slices and arrays can also be referenced directly be pointers
	slice1:= []string{"hello","there","you"}
	var3:=&slice1[1] //This is actually 'var var3 *int = &slice1[1]. So, it is pointing to 2nd value of slice1.
	fmt.Println(var3,*var3)
	//Getting the data-type of the pointer
	fmt.Printf("%T, %T",var3,*var3)
	fmt.Println()
	//To get the memory address by Printf(), use '%p'
	fmt.Printf("%p",var3)
	fmt.Println()
	var4:=&slice1[2]
	fmt.Println(var3,var4) //notice that the difference between these two is '10'
	//So, this tells that GoLang gives 10 byte to every string that is being stored. In case of 'int', it is '4'.

	//Similar referencing can be done on maps
	//The pointers still cannot point to a specific key-value in maps.
	map1:= map[string]int{
		"1st": 91,
		"2nd": 20,
	}
	var6:=&map1
	fmt.Printf("%v, %T",var6,var6)
	fmt.Println()

	//Using struct-pointers
	var var5 *egstruct
	var5=&egstruct{
		name:  "vishu",
		value: 64,
	}
	fmt.Println(var5) //this is showing that var5 is actually pointing to the 'egstruct'
}
type egstruct struct {
	name string
	value int
}