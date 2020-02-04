package main

import "fmt"

func main() {
	var1:=rec1(654)
	var1.eg()
	fmt.Println()
	var2:=rec2{
		name:  "vishu",
		value: 78,
	}
	var2.eg()
}
type rec1 int

type rec2 struct {
	name string
	value float64
}

//In GoLang, it is allowed to create two or more methods with same name but their receiver type must be different.
func (rec rec1) eg() { //we can give the return type here but it is not necessary
	fmt.Println(rec)
}
func (rec rec2) eg() { //we can give the return type here but it is not necessary
	fmt.Println(rec.name,rec.value)
}