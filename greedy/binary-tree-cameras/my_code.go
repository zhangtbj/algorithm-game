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
	arr := []interface{}{0, 0, nil, 0, 0}
	root := BuildTreeFromArray(arr)
	//fmt.Println("Build binary-tree:")

	result = minCameraCover(root)

	fmt.Println(result)
}

const (
	N = 0 //无状态
	C = 1 //摄像头
	Y = 2 //有状态，NULL节点为有状态
)

var result int

func minCameraCover(root *TreeNode) int {
	result = 0
	traversal(root)

	if root.Val == N {
		result++
	}

	return result
}

func traversal(node *TreeNode) int {
	if node == nil {
		return Y
	}

	left := traversal(node.Left)
	right := traversal(node.Right)

	if left == Y && right == Y {
		node.Val = N
		return N
	}
	if left == N || right == N {
		result++
		node.Val = C
		return C
	}
	if left == C || right == C {
		node.Val = Y
		return Y
	}
	return -1
}
