package src_test

import (
	"github.com/stretchr/testify/suite"
	"suitepkg"
	"testing"
)

type StackSuite struct {
	suite.Suite
}

func TestStackSuite(t *testing.T) {
	suite.Run(t,new(StackSuite))
}

func (s *StackSuite) TestEmpty() {
	stack:=suitepkg.NewStack()
}
