package main

import "fmt"

//Embedded structs are the structs in which one struct is utilized in another struct
//So, we get the option to declare the values of a struct from another struct
//This is also known as type-embedding
type College struct {
	Location string
	Fees int
}
type Student struct {
	College
	Name string
	Number int
}
func main() {
	stud:=Student{}
	stud.Location="bhopal"
	stud.Fees=46512
	stud.Name="vishu"
	stud.Number=794
	fmt.Println(stud)
	fmt.Println(stud.Name)
	fmt.Println()

	stud1:=Student{
		College: College{
			Location:"pune",
			Fees:80000,
		},
		Name:    "rishabh",
		Number:  7948,
	}
	fmt.Println(stud1)
	fmt.Println()

	furn:=Furniture{}
	furn.Quality=false
	furn.Time=50
	furn.Item.Name="chair"
	furn.Item.Type="wood"
	fmt.Println(furn)
}
//A struct type can also be declared in a struct
type Furniture struct {
	Time int
	Quality bool
	Item struct{
		Name string
		Type string
	}
}
