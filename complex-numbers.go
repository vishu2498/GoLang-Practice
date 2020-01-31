package main

import "fmt"

func main() {
	//Complex numbers for GoLang are available in two types: complex64 and complex128
	//complex64 refers to float32 for storage and conversion
	//complex128 refers to float64 for storage and conversion
	var var1 complex64 = 3 + 9i
	fmt.Printf("%v, %T",var1,var1)
	fmt.Println()
	var var2 complex128 = 24 + 97i
	fmt.Printf("%v, %T",var2,var2)
	fmt.Println()

	//If we don't define the real part and only provide imaginary part, GoLang will show that real part as '0'
	var var3 complex128 = 53i
	fmt.Printf("%v, %T",var3,var3)
	fmt.Println()
	fmt.Println()

	//Performing operations on complex numbers
	var4:=27+19i
	var5:=65+32i
	add:=var5+var4
	sub:=var5-var4
	mul:=var5*var4
	div:=var5/var4
	fmt.Println(add)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)

	//Getting the real and imaginary part from the complex number
	//real() and imag() refer to float32 while functioning for complex64 number and float64 while functioning for complex128
	var6:=94+61i
	fmt.Println(real(var6))
	fmt.Println(imag(var6))
	fmt.Println()

	//complex() allows us to convert normal numbers into complex numbers
	var7:=complex(42,69)
	fmt.Println(var7)
}
