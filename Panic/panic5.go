package main

import "fmt"

func main() {
	//Runtime panics are those panics that are created automatically and not provided by us
	/*Panics can also be caused by runtime errors such as array out of bounds access.
	This is equivalent to a call of the built-in function panic with an argument defined by interface type 'runtime.Error'
	 */
	runtime()
	fmt.Println("returning from main()")

}
func runtime()  {
	slice:=[]int{64,98,34}
	fmt.Println(slice[5]) //notice that slice[5] doesn't exist
	fmt.Println("slice[5] doesn't exist")
	//running this function will give a panic of index out of range
	//this is how runtime panics get created
}
