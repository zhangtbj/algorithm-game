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

	fmt.Println(removeNthFromEndTest(head, 2))
}

func removeNthFromEndTest(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummyHead.Next
}
