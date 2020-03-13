package suitepkg

type Stack struct {
	isEmpty bool
}

func (stack *Stack) IsEmpty() bool {
	return stack.isEmpty
}

func (stack *Stack) Bury(item string) {
	stack.isEmpty=false
}

func NewStack() *Stack {
	return &Stack{true}
}
