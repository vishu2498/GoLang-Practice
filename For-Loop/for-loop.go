//There is no 'while' loop in GoLang
package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	//Changing the incrementer
	for i := 0; i < 4; i = i + 2 {
		fmt.Print(i)
	}
	fmt.Println()

	//In GoLang, increment operator is a statement and not an expression
	//So, while incrementing multiple variables, 'i++, j++' is not allowed
	//To initialize two or more variables together in for-loop, initialize both of them without declaring each one separately with commas
	for i, j := 0, 1; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	fmt.Println()

	//The initializer variable can be left out
	var1 := 5
	for ; var1 < 10; var1++ {
		fmt.Println(var1)
	}
	fmt.Println()

	//The increment operator can be left out too. So, we have to just define it inside for-loop
	var2 := 5
	for var2 < 10 {
		fmt.Println(var2)
		var2++
	}
	fmt.Println()

	//The above for-loop can also be written as
	var3 := 5
	for var3 < 10 {
		fmt.Println(var3)
		var3++
	}
	fmt.Println()

	//The condition operator can also be left out
	//Although some condition must be written inside for-loop so that it doesn't become infinite
	var4 := 5
	for {
		fmt.Println(var4)
		var4++
		if var4 == 10 {
			break //Using 'break' keyword, the for loop will exit when this condition satisfies
		}
	}
	fmt.Println()

	//Using 'continue' statement
	//'continue' statement will tell the for-loop that whenever the condition satisfies, don't execute the statements inside for-loop and go for next iteration
	for i := 0; i < 15; i++ {
		if i%2 == 0 {
			continue
		}
		//So, we are skipping even values
		fmt.Println(i)
	}
	fmt.Println()

	//Nested for-loop
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i*j)
		}
	}
	fmt.Println()

	//Labels for for-loop are used when we want to break out of the loop due to some condition to a described level.
	//It means that wherever that labels is defined, whenever we use that label, the loop will break at the for-loop.
	Loop: //label is declared where we want the loop to break
		for i:=1;i<=3 ;i++  {
			for j:=1;j<=3 ;j++  {
				fmt.Println(i*j)
				if i*j>=3 {
					break Loop //the name of the loop needs to be provided so that it break from here to that point where it is declared
				}
			}
		}
}
