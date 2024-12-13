package main

type MyQueue struct {
	List []int
}

func (this *MyQueue) Push(x int) {
	this.List = append(this.List, x)
}

func (this *MyQueue) Pop() int {
	first := this.List[0]
	this.List = this.List[1:]
	return first
}

func (this *MyQueue) Peek() int {
	return this.List[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.List) > 0 {
		return false
	} else {
		return true
	}
}

type MyStack struct {
	MyQueue MyQueue
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.MyQueue.Push(x)
}

func (this *MyStack) Pop() int {
	size := len(this.MyQueue.List) - 1
	for i := size; i >= 1; i-- {
		this.MyQueue.Push(this.MyQueue.Pop())
	}
	return this.MyQueue.Pop()
}

func (this *MyStack) Top() int {
	size := len(this.MyQueue.List) - 1
	for i := size; i >= 1; i-- {
		this.MyQueue.Push(this.MyQueue.Pop())
	}

	result := this.MyQueue.Peek()
	this.MyQueue.Push(this.MyQueue.Pop())
	return result
}

func (this *MyStack) Empty() bool {
	return this.MyQueue.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
