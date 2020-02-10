package main

import "fmt"

func main() {
	//Anonymous Struct are the structs which don't have any name and are used immediately after declaring
	//First curly brace of the anonymous struct declares its structure.
	//Second curly brace of the anonymous struct defines its values immediately using initializer syntax.
	struct1:= struct {
		name string
		number int
	}{
		name: "vishu",
		number: 45,
	}
	fmt.Printf("%v, %T",struct1,struct1) //Getting the data-type of variable defined as struct wil return the struct's fields and their data-types.
	fmt.Println()

	//Unlike maps and slices, struct refers to independent data-sets.
	//So, if any struct variable is declared equal to another struct variable, then changing field values from one struct will not change the field values on another struct.
	struct2:= struct {
		name string
		number int
	}{
		"rishabh",
		67,
	}
	struct3:=struct2
	struct3.number=94
	fmt.Println(struct2) //Even when struct3 changed the value, the value of struct2 didn't change
	fmt.Println(struct3)

	//But if we want many struct variables to have the same struct when they are declared equal to each other, we can use '&' literal
	//So, all of the struct variables now point to the same struct and not their copy
	struct4:= struct {
		name string
		number int
	}{
		"associate",
		52,
	}
	struct5:=&struct4
	struct5.number=21
	fmt.Println(struct4) //When the value is changed by struct5, the change can be seen in struct4 too
	fmt.Println(struct5) //Output shows that struct5 points to the same struct of struct4


}
