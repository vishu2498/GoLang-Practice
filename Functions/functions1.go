package main

import "fmt"

//Lowercase functions are only visible to package-level (local-level)
//Uppercase functions are visible everywhere (global-level)
func main() {
	//Every program has the main entry point as main() function
	f1("vishu") //this will print "vishu"

	f2("bruh",15,false)
	//Calling functions in loop
	for i:=1;i<=3 ;i++  {
		f2("hello",51,true)
	}

	f3("1st","2nd")
	//Passing pre-defined variables into the function
	var1,var2:="dell","hp"
	f3(var1,var2)

	msg:="hi"
	value:=45
	f4(msg,value)
	fmt.Println()
	f5(&msg,&value)
}

func f1(msg string)  { //the argument will take the input from the calling function
	fmt.Println(msg) //this will print whatever the calling function has provided in its arguments
}

func f2(msg1 string, value1 int, res1 bool) { //Passing multiple arguments
	fmt.Println(msg1,value1,res1)
}

//If function has multiple arguments that are of the same data-type, there is not need to define the data-type for each of the argument.
func f3(msg,word string)  { //Both arguments are of the same data-type
		fmt.Println(msg,word)
}

//If we change the values of the arguments inside the function, then that value will be permanent for the arguments after it is defined
func f4(msg string,value int)  {
	fmt.Println(msg,value) //this will print whatever will be passed from the arguments of calling function
	value=236 //Hard-coding the argument value
	fmt.Println(msg,value) //output shows that even if a variable holding different value is passed from the calling function twice, hard-coded value will still print after the first execution
	//But, the actual variable value didn't change in the main() function since we are dealing with copy of data.
	//So, in GoLang, arguments are passed by value.
}
func f5(msg *string,value *int)  {
	fmt.Println(*msg,*value)
	*value=512
	fmt.Println(*msg,*value)
	//Now, since we added de-referencing pointers, the changes to actual data will be made and remain permanent.
	//So, no copy of the data will be made.
}