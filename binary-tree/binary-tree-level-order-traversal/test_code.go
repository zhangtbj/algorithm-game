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
	arr := []interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}
	root := BuildTreeFromArray(arr)
	//fmt.Println("Build binary-tree:")

	result := postorderTraversal(root)

	fmt.Println(result)
}

func postorderTraversal(root *TreeNode) []int {
	stack := MyStack{root}
	var result []int

	for !stack.IsEmpty() {
		node := stack.Pop()
		if node == nil {
			continue
		} else {
			result = append(result, node.Val)
		}

		stack.Push(node.Left)
		stack.Push(node.Right)
	}

	start, end := 0, len(result)-1
	for start < end {
		result[start], result[end] = result[end], result[start]
		start++
		end--
	}

	return result
}

type MyStack []*TreeNode

func (stack *MyStack) Push(node *TreeNode) {
	*stack = append(*stack, node)
}

func (stack *MyStack) Pop() *TreeNode {
	if stack.IsEmpty() {
		panic("stack is empty")
	}
	node := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return node
}

func (stack *MyStack) IsEmpty() bool {
	if len(*stack) == 0 {
		return true
	}
	return false
}
