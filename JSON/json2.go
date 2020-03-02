//Encoding JSON
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Marshaling is the process of transforming data into a JSON string.
func main()  {
	map1:=make(map[string]interface{})
	map1["name"]="vishu"
	map1["age"]=21
	map1["alive"]=true

	//Marshal() takes the data of any data-type (since argument type is interface{}) and returns the JSON in the form of slice of bytes.
	//It also returns an error which needs to be dealt with.
	//This is the complete opposite of what Unmarshal() does.
	//So, GoLang objects are converted to JSON or YAML afterwards.
	//While encoding every data into JSON format, Marshal() takes the data and splits them into slice of bytes.
	//Then these slice of bytes are traversed & appended recursively in an ordered manner to generate JSON.
	//Since the final output is slice of bytes only, it needs to be type-casted to another data-type like string.
	output1,err1:=json.Marshal(map1)
	if err1!=nil {
		log.Println("Error is: ",err1)
	}
	fmt.Printf("%T, %v",output1,output1) //output shows that the slice of bytes contain the ASCII codes of the corresponding values (characters)
	fmt.Println()
	fmt.Println(string(output1)) //final required JSON
	//Note that the final JSON output is alphabetically sorted just like maps are alphabetically sorted.
	//This proves that JSON encoding is done from maps.
	fmt.Println()

	//MarshalIndent() works just like Marshal() i.e. it takes the interface object and converts into JSON format.
	//It uses reflection to determine how to transform the map type into a JSON string.
	//But it also allows indenting the final output JSON and adding a new line for every field. Hence, it shows JSON exactly the way it needs to be shown.
	//It needs 3 arguments.
	//1st argument takes the interface object (any data-type) and converts it into JSON.
	//2nd argument is for adding any kind of prefix to the starting of all the fields of JSON in the form of string. This can be left empty too.
	//3rd argument takes how may spaces are needed (in the form of string) for indentation for all the fields from the left. It can also take characters for indentation. This can be left empty too.
	//Similar to Marshal(), it returns a slice of bytes (ASCII codes of actual characters) and an error (which needs to be dealt with).
	output2,err2:=json.MarshalIndent(map1,""," ")
	if err2!=nil {
		log.Println("Error is: ",err2)
	}
	fmt.Println(string(output2))
}