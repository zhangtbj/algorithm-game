package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTreeFromArray(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	// 将根节点初始化
	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root} // 用队列来存储当前层的节点
	index := 1                 // 从数组的第二个元素开始处理

	for index < len(arr) {
		// 取队列头节点作为当前父节点
		current := queue[0]
		queue = queue[1:]

		// 处理左子节点
		if arr[index] != nil {
			leftNode := &TreeNode{Val: arr[index].(int)}
			current.Left = leftNode
			queue = append(queue, leftNode)
		}
		index++

		// 处理右子节点
		if index < len(arr) && arr[index] != nil {
			rightNode := &TreeNode{Val: arr[index].(int)}
			current.Right = rightNode
			queue = append(queue, rightNode)
		}
		index++
	}

	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	arr := []interface{}{3, 2, 3, nil, 3, nil, 1}
	root := BuildTreeFromArray(arr)
	//fmt.Println("Build binary-tree:")

	result := rob(root)

	fmt.Println(result)
}

func rob(root *TreeNode) int {
	result := robTree(root)
	return max(result[0], result[1])
}

func robTree(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}

	leftdp := robTree(node.Left)
	rightdp := robTree(node.Right)

	//不选这个node
	val0 := max(leftdp[0], leftdp[1]) + max(rightdp[0], rightdp[1])

	//选这个node
	val1 := node.Val + leftdp[0] + rightdp[0]

	return []int{val0, val1}
}
