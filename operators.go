package main

import "fmt"

func main() {
	var1:=64
	var2:=32
	result1:=var1+var2
	result2:=var1-var2
	result3:=var1*var2
	result4:=var1/var2
	result5:=var1%var2
	fmt.Println("Addition:",result1)
	fmt.Println("Subtraction:",result2)
	fmt.Println("Multiplication:",result3)
	fmt.Println("Division:",result4)
	fmt.Println("Mod:",result5)
	fmt.Println()

	//Value of Relational Operators are in boolean type
	fmt.Println(var1==var2)
	fmt.Println(var1!=var2)
	fmt.Println(var1>var2)
	fmt.Println(var1<var2)
	fmt.Println(var1>=var2)
	fmt.Println(var1<=var2)
	fmt.Println()

	//Logical Operators
	if(var1!=var2 && var1<=var2){
		fmt.Println("True")
	}

	if(var1!=var2 || var1<=var2){
		fmt.Println("True")
	}

	if(!(var1==var2)){
		fmt.Println("True")
	}

	//Misc Operators
	var3:=17
	var4:=&var3 //var4 is holding the address of var3
	fmt.Println()
	fmt.Println("Address of var3: ",var4)
	fmt.Println("Value of var3: ",*var4)
	//If we want to change the value of var3 via var4, use '*' with var4 and change its value
	*var4=89
	fmt.Println(var3)
	fmt.Println()

	//Bitwise Operators
	// & = AND
	// | = OR
	// ^ = XOR
	// &^ = AND NOT (bit clear operator)
	// << = Left Shift
	// >> = Right Shift

	var5:=10 // Binary: 1010
	var6:=3 //Binary: 0011
	fmt.Println(var5 & var6) // 0010 = 2
	fmt.Println(var5 | var6) // 1011 = 11
	fmt.Println(var5 ^ var6) // 1001 = 9
	fmt.Println(var5 &^ var6) // 1000 = 8
	fmt.Println()
	var7:=8 //2^3
	fmt.Println(var7 << 3) //2^3 * 2^3 = 2^6
	fmt.Println(var7 >> 3) //2^3 / 2^3 = 2^0
}
