package main

import "fmt"

func main() {
	//Struct is essentially a collection of various types of data in variables
	//Struct can be declared both at global and local level
	//Structs are generally the structure of an organized data-type. Variables use them to define their values.
	//Although structs do not have an “empty” value, struct pointers do, with the empty value being nil.
	type Student struct {
		name string
		class int
		number []int
	}
	stud1:=Student{
		name:   "vishu",
		class:  10,
		number: []int{ //when implementing slices in struct, use this syntax
			4651,
			9820,
			1974,
		},
	}
	fmt.Println(stud1)

	//Getting single value from struct
	fmt.Println(stud1.name)
	fmt.Println(stud1.number[1])

	//Its not always needed to write the field names while providing values for struct
	//This is known as positional-syntax. It requires the fields to be in correct order.
	//But the drawback of this method is that we have to define the values in the order they are mentioned in the struct
	//Also, this method can't be used when we are not mentioning every value of fields of struct
	stud2:=Student{
		"rishabh",
		11,
		[]int{
			1234,
			7845,
			197,
		},
	}
	fmt.Println(stud2)
	//While when the field names are defined explicitly, even if they are unordered according to struct, they will work properly.
	//Also with the mention of field names, only certain fields can also be provided.
	stud3:=Student{
		class:12,
		name:"saras",
	}
	fmt.Println(stud3)

	comp1:=Company{
		Name:     "vishu",
		Location: "bhopal",
		Fax:      4658487,
	}
	fmt.Println(comp1)
}
//If we want to use the struct and its fields in another packages, capitalize every first letter of every field of struct
//In case if any of the name of field doesn't start with a capital letter, that field can't be used in another packages/files i.e. it will remain unexported
type Company struct {
	Name string
	Location string
	Fax int
}
