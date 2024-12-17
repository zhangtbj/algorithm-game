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
	arr := []interface{}{1, nil, 2, 3}
	root := BuildTreeFromArray(arr)
	//fmt.Println("Build binary-tree:")

	result := inorderTraversal(root)

	fmt.Println(result)
}

func inorderTraversal(root *TreeNode) []int {
	stack := MyStack{}
	var result []int
	current := root

	for current != nil || !stack.IsEmpty() {
		if current != nil {
			stack.Push(current)
			current = current.Left
		} else {
			current = stack.Pop()
			result = append(result, current.Val)
			current = current.Right
		}
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
