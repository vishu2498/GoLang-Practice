package main

import "fmt"

//Method is just a function that's executing in a known context.
//A known context can be of any data-type like struct, int, map etc.
//A method is a function that contains a receiver argument. This receiver argument can be of any data-type.
//When the method get called, it receives a copy of the receiver argument and uses it via the variable defined with it.
//The concept of receiver argument is that it tells the method that the method must deal with the data-type that's described by the receiver argument.
func main() {
	g:=greeter{
		greeting: "hello",
		name:     "vishu",
	}
	g.greet()
	fmt.Println()

	g1:=greeter{
		greeting: "hello",
		name:     "vishu",
	}
	g1.greet1()
	fmt.Println("New name:",g1.name) //output shows that method implemented the new hard-coded name
}

type greeter struct {
	greeting string
	name string
}

func (g greeter) greet() { //when the receiver argument isn't a pointer, its called a value receiver.
	//So, we are getting a copy of 'greeter' struct and not the actual struct itself.
	fmt.Println(g.greeting,g.name)
}

//However if we want to get the actual struct and be able to manipulate the underlying data, we have to make a pointer receiver.
func(g1 *greeter) greet1() { //when the receiver becomes a pointer, it will point to the actual data from the data-type.
	//So, when we change the value from the method and call this method from the calling function, the new data gets displayed.
	fmt.Println(g1.greeting,g1.name)
	g1.name="rishabh"
}

/*Workflow of the program:
1. We define a data-type like int, struct, map etc.
2. The method contains the receiver as this data-type with a variable so that we can access it.
3. We can use the fields of the data-type in this method by using the receiver variable.
4. This method now gets structure in which it will receive the data.
5. So, we provide the data into this method from any variable from calling function. This data is given to the method in the same structure of its receiver data-type.
6. The methods provides something in result.
7. We get this result via the variable defined equal to the method in the calling function.
 */
