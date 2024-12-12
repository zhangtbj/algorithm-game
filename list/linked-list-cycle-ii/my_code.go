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

	fmt.Println(detectCycle(head))
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast, index1, index2 := head, head, head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			index1 = head
			index2 = fast

			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}

			return index1
		}
	}

	return nil
}
