package main

import "fmt"

func main() {
	var nums = []int{1, 2, 6, 3, 4, 5, 6}
	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	var val = 6

	fmt.Println(removeElementsTest(head, val))
}

func removeElementsTest(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	for head != nil && head.Val == val {
		head = head.Next
	}

	current := head
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
