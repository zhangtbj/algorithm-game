package main

import "fmt"

func main() {
	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3

	fmt.Println(maxSlidingWindow(nums, k))
}

type MyQueue struct {
	K    int
	List []int
}

func (queue *MyQueue) Push(x int) {
	for len(queue.List) != 0 && x > queue.List[len(queue.List)-1] {
		queue.List = queue.List[:len(queue.List)-1]
	}
	queue.List = append(queue.List, x)
}

func (queue *MyQueue) Pop(x int) {
	if len(queue.List) > 0 && queue.List[0] == x {
		queue.List = queue.List[1:]
	}
}

func (queue *MyQueue) Peek() int {
	return queue.List[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int

	myQueue := MyQueue{K: k}
	for i := 0; i < k; i++ {
		myQueue.Push(nums[i])
	}
	result = append(result, myQueue.Peek())

	for i := k; i < len(nums); i++ {
		myQueue.Pop(nums[i-k])
		myQueue.Push(nums[i])
		result = append(result, myQueue.Peek())
	}
	return result
}
