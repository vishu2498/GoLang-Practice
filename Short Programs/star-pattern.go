package main

import "fmt"

func main() {
	for i:=0;i<5;i++{
		for j:=0;j<=i;j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	for i:=5;i>0;i-- {
		for j:=i; j>0;j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	for i := 1; i <= 3; i++ {
		for j := i; j < 3; j++ {
			fmt.Print(" ")
		}
		for k := 1; k < (i * 2); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 3; i >= 1; i-- {
		for j := 3; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k < (i * 2); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	for i:=1;i<=5;i++{
		for j:=4;j>i-1;j--{
			fmt.Print(" ")
		}
		for k:=1; k<=i; k++{
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
