package main

import (
	"fmt"
	"reflect"
)
type order struct {
	ordId      int
	customerId int
}
func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)
}
func createQuery1(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}
func main() {
	/* Reflection gives you the ability to examine types at runtime.
	It also allows you to examine, modify, and create variables, functions, and structs at runtime.
	The reflect package helps to identify the underlying concrete type and the value of a interface{} variable.
	*/
	var1 := 567
	var var2 interface{}
	result := reflect.TypeOf(var1)
	fmt.Println("Value of var1 is: ", var1)
	fmt.Println("Data type of var1 is: ", result)
	fmt.Println("Type of var2 interface: ", reflect.TypeOf(var2))
	fmt.Println("Value of var2 interface: ", reflect.ValueOf(var2))
	fmt.Println("Getting kind, type & values of a struct using reflect package: ")
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)
	fmt.Println("Getting field numbers: ")
	createQuery1(o)
	var3 := 89
	x := reflect.ValueOf(var3).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	var4 := "Vishu"
	y := reflect.ValueOf(var4).String()
	fmt.Printf("type:%T value:%v\n", y, y)
	fmt.Println("Copying data from one slice to another: ")
	src := reflect.ValueOf([]int{1, 2, 3})
	/* make sure the dest space is larger than src */
	dest := reflect.ValueOf([]int{10, 20, 30})
	cnt := reflect.Copy(dest, src)
	fmt.Println(cnt)
	fmt.Println(src)
	fmt.Println("Using different data types from reflect: ")
	for _, v := range []interface{}{"vishu", 98,} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println("String is: ", v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println("Value is: ", v.Int())
		}
	}
}