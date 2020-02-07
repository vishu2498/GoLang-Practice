package main

import "fmt"

func main() {
	//Slices can also be made by make() function
	//1st argument will be the data-type of slice, 2nd will be length of slice and 3rd will be capacity of slice
	//Length of the slice can't be larger than capacity of slice. If we do so, compiler will give an error: "len larger than cap in make([]int)".
	slice1:=make([]int,3,5) //It is possible to have different length and capacity of a slice (unlike arrays)
	fmt.Println("Length:",len(slice1))
	fmt.Println("Capacity:",cap(slice1))
	//But when we print the slice having different length and capacity, it will only display the elements with the length given
	fmt.Println(slice1)
	fmt.Println()

	//Adding elements to a slice using append() function
	//append() function is actually a variadic function
	//1st argument is name of the slice and further arguments are the elements which we want to add
	slice2:=[]int{2,3,5}
	slice2=append(slice2,6,7,8)
	fmt.Println(slice2)
	/*It is important to note that when elements are appended to the slice and if the slice doesn't have capacity explicitly defined,
	the capacity of the slice now becomes approximately '2*length'. In case if any element is appended to the slice which doesn't have any values (zero length and capacity),
	GoLang compiler will first assign a memory location to the slice and then copy the elements that were first present in the slice.
	Since any previous elements were not present, compiler will then append the given values and thereby making of capacity
	of the slice as approximately '2*length'. However, it always not confirmed that how compiler will assign the capacity of
	the same or new slice.
	 */
	fmt.Println()

	//Slice-Concatenation (using append() function)
	//While adding one slice to another slice, the data-type of both must be the same
	slice3:=[]int{5,1,9}
	slice3=append(slice3, []int{2,5,8}...) //Using '...' literal here is extremely important. It tells GoLang compiler that just break the elements given in slice into individual elements of the given data-type and append the values to the slice
	//Above append() is exactly equal to writing: "slice3=append(slice3, 2, 5, 8)
	fmt.Println(slice3)
	fmt.Println()

	//Slicing slices and storing as another slice
	slice4:=[]int{2,3,4,8,9}
	slice5:=slice4[1:] //slice5 will not have the 1st element of slice4
	fmt.Println(slice5)
	slice6:=slice4[:len(slice4)-2] //slice6 will not have the last 2 elements of slice4
	fmt.Println(slice6)
	//Slicing the slice from the middle
	slice7:=append(slice4[:2],slice4[3:]...) //this is keeping the first 2 elements and last 2 elements of slice4
	fmt.Println(slice7)
	fmt.Println(slice4) //Notice that the values of slice4 are changed due to the operation performed by slice7

	//When we build a slice with make() function by giving its length and capacity and slice that slice with three index slice, the new slice has again reduced length and capacity.
	slice8:=[]string{"1st","2nd","3rd","4th","5th"}
	slice9:=slice8[2:4:5]
	//1st number represents to refer to 3rd element of slice8, 2nd number represents till where to slice from the 1st number position, 3rd number defines the capacity of slice but should not be lower than 2nd number and should not be more than the range of original slice.
	fmt.Println(slice9)
}
