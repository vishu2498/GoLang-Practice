package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file3, err := os.Create("file3.txt")
	if err != nil {
		file3.Close()
		log.Print("couldn't create the file")
		return
	}

	slice1 := []string{"message from vishu", "how are you", "one minute... what?"}

	for _, values := range slice1 {
		//'fmt.Fprintln()' takes 1st argument as "io.Writer" and writes the contents into it.
		//It takes the contents to be written from the 2nd argument onwards (it's a variadic function).
		fmt.Fprintln(file3, values) //writing elements of slice one by one into the file
		if err != nil {
			log.Println(err)
			return
		}
	}

	errclose3 := file3.Close()
	if errclose3 != nil {
		log.Fatal("couldn't close the file")
	}
}
