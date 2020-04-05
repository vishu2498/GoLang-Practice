package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Showing that not every field in struct is added for JSON. Hence, those neglected (not-added) fields will not appear in the final result.
//This gives the option and flexibility to get JSON of any number of fields and not using every field.
type Details struct {
	Name string `json:"name"`
}

var JSON = `{
	"name": "vishu",
	"age": 22
}`

func main() {
	var data Details
	err := json.Unmarshal([]byte(JSON), &data)
	if err != nil {
		log.Fatal("error received")
	}
	fmt.Println(data) //Even if two fields are present in the JSON, only one field is displayed which is mentioned in the struct.
}
