//Decoding JSON
package main

//JSON package (encoding/json) is used for performing operations on JSON files and its contents.
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//This struct contains the fields corresponding to the structure of the JSON.
//Every field needs to be exported in order for JSON to be executed.
//Also, every field should have the tag of 'json' to tell the GoLang compiler that which JSON field is attached to which struct field.
type Result struct {
	Subject string `json:"subject"`
	Marks int `json:"marks"`
	Verdict bool `json:"verdict"`
}

//Providing JSON locally
//Since variable is at package-level, short declaration operator (:=) can't be used.
var JSON1 = `{
	"subject": "maths",
	"marks": 68,
	"verdict": true
}`

//Using JSON array
var JSON2 = `[
{
	"name": "vishu",
	"designation": "intern",
	"exp": 8
}
]`

//JSON with multiple inside fields
var JSON3 = `{
	"name": "vishu",
	"age": 21,
	"contact": {
		"number": 23124541,
		"email": "example@eg.com"
	}
}`

func main() {
	var final Result

	//Unmarshal() converts the data from the JSON value to an object of interface type. Since the object is of interface type, the object can have any data-type.
	//It takes 2 arguments. 1st argument is the data (generally string) which is converted into slice of bytes. 2nd argument is the referenced variable which is of any data-type representing the final output.
	//Unmarshal() matches incoming object keys to the keys used by Marshal() (either the struct field name or its tag), preferring an exact match but also accepting a case-insensitive match.
	//This function can also be used to convert YAML or JSON into a GoLang object (generally a map). This is done by 1st converting the YAML to JSON and then converting JSON to GoLang object.
	err1:=json.Unmarshal([]byte(JSON1),&final)
	if err1!=nil {
		log.Println("Error is :",err1)
		return
	}
	fmt.Println(final) //printing the values of JSON
	fmt.Printf("%T",&final) //showing that this referenced variable is of type pointer
	fmt.Println()
	fmt.Printf("%T",JSON1) //showing that entire JSON is string
	fmt.Println()
	fmt.Println()

	//showing that JSON unmarshalling doesn't always need a struct or any other data-type
	//Since Unmarshal() generally returns a map, a custom map can be created which will store the values returned by the Unmarshal() function.
	//Keeping the 2nd data-type of map as interface{} since value can be of any data-type.
	var finalres map[string]interface{}
	err3:=json.Unmarshal([]byte(JSON1),&finalres)
	if err3!=nil {
		log.Println("Error is:",err3)
		return
	}
	fmt.Println(finalres["verdict"]) //taking single value from the JSON

	//Unmarshalling JSON array
	var emp []interface{} //taking interface array, without using additional struct for JSON unmarshalling
	err2:=json.Unmarshal([]byte(JSON2),&emp)
	if err2!=nil {
		log.Println("Error is:",err2)
		return
	}
	fmt.Println(emp) //output also shows that the result object is a map
	fmt.Println()

	//Unmarshalling JSON from external file
	//ioutil.ReadFile() function is used to read the contents from file. It returns the contents in slice of bytes. It also returns an errors which needs to be dealt with.
	externalfile,err4:=ioutil.ReadFile("C:/Users/vishwanath.taykhande/Desktop/GoLang-Practice/JSON/eg.json")
	var ext interface{} //taking an interface{} so that it can store any kind of data
	err5:=json.Unmarshal([]byte(externalfile),&ext)
	if err4!=nil {
		log.Println("Error is:",err4)
		return
	}
	if err5!=nil {
		log.Println("Error is:",err5)
		return
	}
	fmt.Println(ext)
	fmt.Println()

	//Getting individual values from the JSON
	var multijson map[string]interface{}
	err6:=json.Unmarshal([]byte(JSON3), &multijson)
	if err6!=nil {
		log.Println("Error is:",err6)
		return
	}
	fmt.Println("Name:",multijson["name"])
	//When values are inside the fields in JSON, they have to be first type-casted to the output type and called from their parent field.
	fmt.Println("Email:",multijson["contact"].(map[string]interface{})["email"])
}