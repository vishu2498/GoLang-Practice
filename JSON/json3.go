package main

import (
	"encoding/json"
	"fmt"
)

//Using JSON "omitempty" tag
//This is particularly used for encoding JSON.

//JSON "omitempty" tag is used when we don't know the value of a field and also don't want to represent it with its default or zero value.
//So, using this tag, the field can be neglected and can be completely removed from the output.
type emp struct {
	Name string `json:"name"`
	Age int `json:",omitempty"` //this syntax tells compiler that this JSON field can be neglected
	Salary *int `json:",omitempty"` //this syntax tells compiler that the field is of 'int' pointer and can only be omitted when there is no value given to it
	Address *[]string `json:",omitempty"` //intentionally adding this field by providing no value with it
}

func main() {
	var sal int=0 //providing zero/default value of 'int' in a variable
	emp1:=emp{
		Name: "vishu",
		Age: 0, //since we provided zero value of 'int' here, this field got neglected in the output
		//The same will happen if a string is empty "", or if a pointer is nil, or if a slice has zero elements in it.

		//However, since 'Salary' field is an 'int' pointer with 'omitempty' tag, we need to send a dereferenced pointer to the variable holding the data that we want to represent as the field's value.
		//So, this allows us to send the zero values too to the fields with 'omitempty' tags without neglecting them in output.
		Salary: &sal,
	}
	result1,_:=json.Marshal(emp1)
	fmt.Println(string(result1))
}
