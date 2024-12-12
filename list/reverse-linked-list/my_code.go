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

	fmt.Println(reverseList(head))
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	for cur != nil {
		temp := cur.Next
		pre = cur
		cur = temp
		cur.Next = pre
	}

	return pre
}
