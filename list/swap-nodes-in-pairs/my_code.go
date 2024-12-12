package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var nums = []int{1, 2, 3, 4, 5}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		newNode := &ListNode{Val: nums[i]}
		current.Next = newNode
		current = current.Next
	}

	fmt.Println(swapPairs(head))
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	current := dummyHead

	for current.Next != nil && current.Next.Next != nil {
		temp := current.Next
		temp1 := current.Next.Next.Next

		current.Next = current.Next.Next
		current.Next.Next = temp
		temp.Next = temp1

		current = current.Next.Next
	}

	return dummyHead.Next
}
