package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDepth int
var result int

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
	arr := []interface{}{0}
	root := BuildTreeFromArray(arr)
	//fmt.Println("Build binary-tree:")

	result := findBottomLeftValue(root)

	fmt.Println(result)
}

func findBottomLeftValue(root *TreeNode) int {
	maxDepth = -1
	result = 0
	compare(root, 1)
	return result
}

func compare(node *TreeNode, depth int) {
	if node.Left == nil && node.Right == nil {
		if depth > maxDepth {
			maxDepth = depth
			result = node.Val
		}
	}

	if node.Left != nil {
		depth++
		compare(node.Left, depth)
		depth--
	}

	if node.Right != nil {
		depth++
		compare(node.Right, depth)
		depth--
	}
}
