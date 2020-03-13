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
	//Instead of using 'stack:= suitepkg.NewStack()' in every test, we can configure it once here.
	s.stack=suitepkg1.NewStack()
}

func (s *StackSuite) TestEmpty() {
	s.True(s.stack.IsEmpty())
}