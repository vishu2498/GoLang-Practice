package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//'os.Create()' is used to create a new file with the name given or replace it if that file already exists.
	//It takes the name of file in it's argument.
	//The new file is created with mode 0666.
	//In it's internal operation, it calls 'syscall.O_CREAT' which actually creates the files.
	file1,err:=os.Create("file1.txt")
	if err!=nil {
		log.Fatal("couldn't create the file")
		return
	}
	fmt.Printf("%v,%T",file1,file1)
	fmt.Println()
	fmt.Println(file1.Name()) //shows the name of file
	fmt.Print(file1.Stat()) //show the stats and details of the file
	fmt.Println()

	slice1:=[]byte{53,68,24}
	content1,err:=file1.Write(slice1) //It writes contents to the file in the form of slice of bytes and returns the number of bytes written.
	if err!=nil {
		log.Fatal("couldn't write to file")
		file1.Close()
		return
	}
	fmt.Println(content1) //shows the length of the contents written to file
	errclose1:=file1.Close()
	if errclose1!=nil {
		log.Fatal("couldn't close the file")
	}


	file2,err:=os.Create("file2.txt")
	if err!=nil {
		fmt.Println("couldn't create the file")
		return
	}

	content2,err:=file2.WriteString("message from vishu") //It is similar to 'Write()' but directly takes string as argument instead of slice of bytes.
	if err!=nil {
		fmt.Println("couldn't write to file")
		file2.Close()
		return
	}
	fmt.Println(content2) //length of contents written in the file
	errclose2:=file2.Close()
	if errclose2!=nil {
		log.Fatal("couldn't close the file")
	}
}
