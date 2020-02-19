//Please include GOPATH for execution of this file

package main

import (
	"fmt"
	usf "usfield"
)

func main()  {
	var1:=usf.Employer{
		Dismiss:false,
	}

	//Now, even if the fields of another struct that is embedded but actually is unexported, the struct variable can still access it from another package/file.
	//This is done by first referencing a variable with the struct that embeds another struct and if the fields of the unexported struct are exported, then those fields are accessible.
	var1.Name="vishu"
	var1.Number=652
	fmt.Println(var1)
}
