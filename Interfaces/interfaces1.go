package main

import "fmt"
/* Interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract,
so we are not allowed to create an instance of the interface.
But we are allowed to create a variable of an interface type and this variable can be assigned with a type value that has
the methods the interface requires. Or in other words, the interface is a collection of methods as well as it is a custom type.
 */
//InGolang, interfaces are implemented implicitly.
//All the methods defined inside an interface need to be implemented.
func main() {
	//Default value of interface is '<nil>'
	type Counter interface{} //If there are no methods inside interface, it is called empty interface. So, all types implement the empty interface.
	var cnt Counter
	fmt.Printf("%v, %T", cnt, cnt)
	fmt.Println()

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello vishu"))
}

//Interface doesn't describe data. Instead it describes behaviours.
//So, instead of defining data-types here like struct, methods are declared here.
type Writer interface {
	Write([]byte) (int,error)
}

/*The user-defined type which the method uses as receiver is also called concrete type since interface values have
no concrete behavior without the implementation of the stored user-defined value.*/
type ConsoleWriter struct { //this struct implements Writer interface

}

func (cw ConsoleWriter) Write(data []byte) (int,error) {
	n,err:=fmt.Println(string(data))
	return n,err
}

//It is also not important to describe the argument with the method exactly like that's defined previously. Instead, the fields of struct can directly be written as arguments of the method.
//This allows the method to work directly with the struct fields as only and later these arguments work as the input for the fields of the struct.
/*Workflow of the program:
1. We define an interface that contains a method. This method takes slice of bytes as argument and returns int and error.
2. We define a struct that does not hold any fields.
3. We define the method that was declared in the interface with the same name.
4. This method takes the argument as the struct which was previously defined with its own variable.
5. In the main function, we define a variable whose data-type is the interface we have implemented. This shows that interface is a custom type and can be used to denote any kind of data and can be used by anything.
6. This variable is equal to the struct which was defined previously.
7. Now, since the variable is an interface that can store and use 'struct' type data, the variable is allowed to use the method defined in the interface and pass the data as argument.
8. This data is of slice bytes and the method converts that slice of bytes into string and returns it.
 */
//This shows that if we want to feed the data (by a method) in the data-type that someone else has defined, we can use interface to use and hold that custom type.

