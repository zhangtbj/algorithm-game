package main

// This is the Stack

type Stack struct {
	List []int
}

func (stack *Stack) Push(x int) {
	stack.List = append(stack.List, x)
}

func (stack *Stack) Pop() int {
	first := stack.List[len(stack.List)-1]
	stack.List = stack.List[:len(stack.List)-1]
	return first
}

func (stack *Stack) Peek() int {
	return stack.List[len(stack.List)-1]
}

func (stack *Stack) Size() int {
	return len(stack.List)
}

func (stack *Stack) Empty() bool {
	return len(stack.List) <= 0
}

// This is the Queue

type MyQueue struct {
	StackIn  Stack
	StackOut Stack
}

func Constructor() MyQueue {
	return MyQueue{
		StackIn:  Stack{},
		StackOut: Stack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.StackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	result := this.Peek()
	this.StackOut.Pop()
	return result
}

func (this *MyQueue) Peek() int {
	if this.StackOut.Empty() {
		for !this.StackIn.Empty() {
			this.StackOut.Push(this.StackIn.Pop())
		}
	}
	return this.StackOut.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.StackIn.Empty() && this.StackOut.Empty()
}
