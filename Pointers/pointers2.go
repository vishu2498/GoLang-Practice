package main
import "fmt"

func main() {
	//Default value of a pointer is '<nil>'
	var var1 *int
	fmt.Println(var1) //showing default value of pointer
	//Even if the pointer variable is defined equal to a struct, it will hold a '<nil>' value

	var var2 *struct1
	fmt.Println(var2)

	//Using new() function to to define a new pointer variable for the struct
	//new() function is used to initialize a variable to a pointer to an object
	//However, new() function doesn't allow initializing values
	var3:=new(struct1)
	fmt.Printf("%T, %v",var3,var3)
	//output shows that variable is pointer to the struct and since new() doesn't allow initialization of values, the default value is '0' (due to 'int')
	fmt.Println()

	//De-referencing operator actually has low preference that '.' operator (that is used for accessing functions or variables)
	//To add values to the struct via the pointer variable made with new() function, de-referencing operator can be used
	(*var3).value=321
	fmt.Println(var3) //this is equal to 'fmt.Println((*var3).value)
	//However, de-referencing of pointer variable isn't actually needed to access or modify the values of the struct
	var4:=new(struct1)
	var4.value=251 //this is equal to '(*var4).value=251'
	/*This is made possible by GoLang compiler since it knows that we are trying to modify the value that is present in a struct via the pointer variable.
	So, it will directly attach us to the object that is holding the actual field of the struct. Hence, de-referencing is not needed.*/
}
type struct1 struct {
	value int
}
