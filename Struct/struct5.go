package main
import "fmt"

//Implementing the same struct within struct
type st1 struct {
	f1 string
	f2 int
	st1 struct { //struct of same name redefined
		f3 string
		f4 int
	}
}
func main() {
	var1:=st1{
		f1: "vishu",
		f2: 32,
	}

	//While providing values for the embedded similar name struct, the fields need to be redeclared with values providing afterwards.
	var2:=st1{st1: struct {
		f3 string
		f4 int
	}{f3: "new", f4:52 }}

	fmt.Println(var1)
	fmt.Println(var2)
}