package main

import "fmt"

func main() {
	//recover is like the 'catch' of try-catch concept
	//recover is a builtin function which is used to regain control of a panicking goroutine.
	/*Recover is useful only when called inside deferred functions.
	Executing a call to recover inside a deferred function stops the panicking sequence by restoring normal execution
	and retrieves the error value passed to the call of panic.
	If recover is called outside the deferred function, it will not stop a panicking sequence.
	 */
	/*The recover function returns the value passed to panic function and has no side effects.
	That means if our goroutine is not panicking, recover function will return nil.
	So checking the return value of recover() to nil is a good way to know if your program is packing,
	unless some idiot passes nil to the panic function which is a very rare case.
	 */
	/*Since the only functions which will get executed when the panic occurs are deferred function,
	we need to use recover function inside one of them.
	 */
	/*If a function call is deferred from another deferred function, you won’t be able to recover your program there.
	Because in order to recover a panicking program successfully, a deferred function with recover call must be executed right after the panic.
	 */
	/* Steps for recover function to work:
	1. When any panic occurs from a statement or a function, all the statements after it do not get executed.
	2. Then this panic searches for recover function in the same function and executes it to recover from panic & resumes
	   program execution.
	3. If it doesn't find any recover function inside same function, it searches the function that is holding the
	   recover function in itself and recovers from panic. This function which contains the recover function needs to be
	   deferred in the current function.
	 */
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
func recoverName() {
	if r := recover(); r!= nil {
		fmt.Println("recovered from ", r)
	}
}
func fullName(firstName *string, lastName *string) {
	//when the panic from the if conditions arrives, it moves recoverName function
	//the recoverName function helps in recovery from panic
	//so the program execution continues
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
