package src_test

import (
	"github.com/stretchr/testify/suite"
	"suitepkg"
	"testing"
)

//Suite is defined by its struct.
type StackSuite struct {
	//this needs to be implemented in the struct
	suite.Suite //'Suite' is a struct with pointers to 'assert', 'require' and 'testing' package.
}

//Atleast one function like this needs to be implemented for using 'Run()' and 'go test'.
func TestStackSuite(t *testing.T) {
	suite.Run(t,new(StackSuite)) //this function is required to run the associated tests.
	//It takes the testing variable and an object to the struct made for suite.
}

//From here on, we can write our test-cases.
func (s *StackSuite) TestEmpty() {
	stack:=suitepkg.NewStack()
	s.True(stack.IsEmpty())
}

func (s *StackSuite) TestNotEmpty()  {
	stack:=suitepkg.NewStack()
	stack.Bury("green")
	s.False(stack.IsEmpty()) //using 's.False()' to pass the test
}