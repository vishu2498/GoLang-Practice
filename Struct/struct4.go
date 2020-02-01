package main

import (
	"fmt"
	"reflect" //reflect package is required to use tags in structs
)
//Tags can be used with fields of struct for additional information
type Glasses struct {
	Type string `limited types:"Windshield,Mirror"`
	Quality bool
	Lens int `lens:concave,convex`
}
func main() {
	glass:=reflect.TypeOf(Glasses{
		Type:    "Windshield",
		Quality: true,
		Lens:    2,
	})
	//Getting tags from fields
	tag,_:=glass.FieldByName("Type")
	tag1,_:=glass.FieldByName("Lens")
	fmt.Println(tag.Tag)
	fmt.Println(tag1.Tag)
}
