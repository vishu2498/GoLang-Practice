package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

var JSON = `{
	"name": "vishu",
	"age": 21
}`

func main() {
	fmt.Println(result(JSON))
}

//Decoding the JSON using 'json.NewDecoder()' & 'Decode()' functions.
//It will return the same output as JSON Unmarshalling (JSON as map).
func result(str string) (interface{},error) {
	var result interface{} //this interface will hold the final result
	decodeIntoObj:=json.NewDecoder(strings.NewReader(str)) //variable will hold the '*Decoder' object after decoding from the string
	//The string is needed to be converted to 'Reader' object since 'json.NewDecoder()' needs that data-type in its argument.

	//'Decode()' reads the next JSON-encoded value from its input and stores it in the interface.
	if err:=decodeIntoObj.Decode(&result);err!=nil {
		return nil,err
	}
	return result,nil
}
