package main

import (
	"fmt"
	"time"
)

func main() {
	//Even with the concept of closures, race-condition can arrive.
	//A race condition is when two or more GoRoutines have access to the same resource, such as a variable or data structure and attempt to read and write to that resource without any regard to the other routines.
	//Here, the race condition has occurred between main() GoRoutine and the Anonymous GoRoutine assigned inside it.

	/*So, even if we are defining the variable 2nd time after the initialization of GoRoutine,
	Go-runtime will first fetch and interpret all the variables and its changes outside the GoRoutine.
	This will happen because the Go-runtime scheduler will not interrupt the main() GoRoutine till it hits the time.Sleep() function.
	Hence, even if it launches the GoRoutine first before the 2nd definition of the variable, it will be still in the main() GoRoutine and will first change the value of the variable.
	Then will move on to the execution of GoRoutine.
	*/
	var msg = "hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "bye" //This will be executed before the GoRoutine launches and this value will be passed to the GoRoutine.
	//Hence creating the race condition.
	time.Sleep(100 * time.Millisecond)

	//To solve this race condition, we will pass the variable as an argument to the GoRoutine.
	/*So, GoRuntime is first copying the string "hello" from var1 to the function via its argument and then
	it will decouple the variable in the main() function from the GoRoutine. So, the message being printed by GoRoutine
	is actually a copy of the data first present in the variable. Since the variable is decoupled, now even if it is
	changed after the GoRoutine, it makes no difference since that GoRoutine has now received the data already.
	This is preventing the race condition caused due to dual-assignment of variables.
	 */

	var msg1 = "hello"
	go func(msg1 string) {
		fmt.Println(msg1)
	}(msg1)
	msg1 = "bye"
	time.Sleep(100 * time.Millisecond)
}
