package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Name string
	Age int
}

var db map[string]Student

func addStudent(s Student) error {
	if s.Age >30 {
		return errors.New("Student too old.")
	}
	//If a student's name is already present, error indicates duplicate entry
	_,ok:=db[s.Name]
	if !ok {
		return errors.New("duplicate entry")
	}
	//storing the Student credentials
	db[s.Name]=s
	return nil
}
func main() {
	db=make(map[string]Student)
	s1:=Student{"vishu",35}
	err:=addStudent(s1)
	fmt.Println("Received an error: ",err.Error())

	s2:=Student{"yash",21}
	err=addStudent(s2)
	if err != nil {
		fmt.Println("Adding yash -  should get an empty error")
	}
	err=addStudent(s2)
	fmt.Println("Adding yash twice, got an error",err.Error())
}