package main

import "fmt"

func main() {
	fmt.Println("sum of the given values is:",f1(2,6,8))
	fmt.Println()
	var1:=f2(2,6,8,4,1)
	fmt.Println("sum is:",*var1)
	fmt.Println()
	var2:=f3(2,6,8,4,1)
	fmt.Println("sum is:",var2)
	var3,err:=f5(5,0)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(var3)
}

//Setting the return data-type for function
func f1(values ...int) int{
	result:=0
	for _,v:=range values {
		result = result + v
	}
	//'return' keyword becomes compulsory here
	//The return value must be of the same data-type as the function is returning.
	return result
}

//GoLang can also return a variable as a pointer
/* In every other language, whenever a function block executes, the compiler allocates a memory stack for it.
Now, when that function operation is completed, compiler deletes that local memory stack for freeing up the memory.
So, if any pointer is referencing to anything in that memory stack, that data is now wiped out and the pointer becomes a dangling pointer.
However, in GoLang, when the compiler gets to know that a return statement or pointer variable is pointing to something in a memory stack which will be depleted after function execution,
the GoLang compiler will keep that specific stack into the shared memory of the PC. This memory will be a heap memory.
So, there will be no issue in getting the data from return statement single/multiple times even if the function's execution is done.
 */
func f2(values ...int) *int {
	result1 := 0
	for _, v := range values {
		result1 = result1 + v
	}
	return &result1
}

//Getting named return from a function
func f3(values ...int) (result int) { //Last argument of the function with/without brackets will be its return type
	//Hence, that return variable will be available in the function to use it.
	//Also, we don't have to specify the name of the data/variable with return at the end of the function.
	for _,v:=range values{
		result=result+v
	}
	return
}

//Getting multiple return types from function
func f5(var1,var2 int) (int,error) { //mentioning multiple return types
	//If a function is returning multiple values, then the calling function must use it with as many number of variables as there are return types
	if var2==0 {
		return 0,fmt.Errorf("Cannot divide by zero")
	}
	return var1/var2, nil
}