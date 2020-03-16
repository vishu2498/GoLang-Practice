package suitepkg1_test

import (
	"github.com/stretchr/testify/suite"
	"suitepkg1"
	"testing"
)

type StackSuite struct {
	suite.Suite
	stack *suitepkg1.Stack //added this for implementing 'SetupTest()'
}

func TestStackSuite(t *testing.T) {
	suite.Run(t,new(StackSuite))
}

//This function runs before each test and is called by testify.
//It can also be regarded as a configuration for all the test (like a permanent receiver for all tests).
func (s *StackSuite) SetupTest() {
	//Instead of using 'stack:= suitepkg1.NewStack()' in every test, we can configure it once here.
	s.stack=suitepkg1.NewStack()
}

func (s *StackSuite) TestEmpty() {
	s.True(s.stack.IsEmpty())
}

func (s *StackSuite) TestUnburySingleItem() {
	s.stack.Bury("green")
	s.Equal("green",s.stack.Unbury())
	s.Zero(s.stack.Size())
	s.True(s.stack.IsEmpty())
}

func (s *StackSuite) TestUnburyMultipleItems() {
	s.stack.Bury("green")
	s.stack.Bury("yellow")

	s.Equal("yellow",s.stack.Unbury())
	s.Equal(1,s.stack.Size())

	s.Equal("green",s.stack.Unbury()) //If we don't 'Unbury()' the times 'Bury()' was used, test will fail.
	s.Zero(s.stack.Size())
	s.True(s.stack.IsEmpty())
}