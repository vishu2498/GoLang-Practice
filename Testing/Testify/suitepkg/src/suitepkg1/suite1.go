package suitepkg1

type Stack struct {
	isEmpty bool
	size int
	items [6]string
}

func (stack *Stack) IsEmpty() bool {
	return stack.isEmpty
}

func (stack *Stack) Bury(item string) {
	stack.items[stack.size]=item
	stack.size++
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Unbury() string {
	stack.size--
	return stack.items[stack.size]
}
func NewStack() *Stack {
	return &Stack{true,0,[6]string{}}
}
