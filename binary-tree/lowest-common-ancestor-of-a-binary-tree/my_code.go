package main

import "fmt"

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
	arr := []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	root := BuildTreeFromArray(arr)
	//fmt.Println("Build binary-tree:")

	left := root.Left
	right := root.Right
	result := lowestCommonAncestor(root, left, right)

	fmt.Println(result)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	leftAncestor := lowestCommonAncestor(root.Left, p, q)
	rightAncestor := lowestCommonAncestor(root.Right, p, q)

	if leftAncestor != nil && rightAncestor != nil {
		return root
	} else if leftAncestor != nil && rightAncestor == nil {
		return leftAncestor
	} else if leftAncestor == nil && rightAncestor != nil {
		return rightAncestor
	} else {
		return nil
	}
}
