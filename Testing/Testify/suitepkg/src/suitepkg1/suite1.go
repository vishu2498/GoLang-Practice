package suitepkg1

type Stack struct {
	isEmpty bool
	size int
}

func (stack *Stack) IsEmpty() bool {
	return stack.isEmpty
}

func (stack *Stack) Bury(item string) {
	stack.size++
}

func (stack *Stack) Size() int {
	return stack.size
}
func NewStack() *Stack {
	return &Stack{true,3}
}
