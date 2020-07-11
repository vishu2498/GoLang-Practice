package main

import "fmt"

func main()  {
	cities:=[]string{"New York","New Jersey","Paris","Delhi","Mumbai","Bhopal","Surat","Los Angeles"}
	for i:=len(cities)-1; i>=0; i-- {
		fmt.Print(cities[i]+" ")
	}
}
