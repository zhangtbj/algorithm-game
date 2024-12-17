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
	arr := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := BuildTreeFromArray(arr)
	//fmt.Println("Build binary-tree:")

	result := levelOrder(root)

	fmt.Println(result)
}

func levelOrder(root *TreeNode) [][]int {
	var nodes []*TreeNode
	var result [][]int
	var size = 1

	if root == nil {
		return result
	} else {
		nodes = append(nodes, root)
	}

	for len(nodes) > 0 {
		var subResult []int
		var subSize int
		for size > 0 {
			pop := nodes[0]
			nodes = nodes[1:]
			size--
			subResult = append(subResult, pop.Val)

			if pop.Left != nil {
				nodes = append(nodes, pop.Left)
				subSize++
			}
			if pop.Right != nil {
				nodes = append(nodes, pop.Right)
				subSize++
			}
		}
		result = append(result, subResult)
		size += subSize
	}

	return result
}
