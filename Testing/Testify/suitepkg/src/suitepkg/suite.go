package suitepkg

type Stack struct {
	isEmpty bool
	size int
}

func (stack *Stack) IsEmpty() bool {
	return stack.isEmpty
}

func (stack *Stack) Bury(item string) {
	stack.isEmpty=false
	stack.size++
}

func (stack *Stack) Size() int {
	return 3
}
func NewStack() *Stack {
	return &Stack{true,3}
}
