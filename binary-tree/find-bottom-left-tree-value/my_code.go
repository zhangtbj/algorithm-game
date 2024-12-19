package main

//import "fmt"
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func BuildTreeFromArray(arr []interface{}) *TreeNode {
//	if len(arr) == 0 {
//		return nil
//	}
//
//	// 将根节点初始化
//	root := &TreeNode{Val: arr[0].(int)}
//	queue := []*TreeNode{root} // 用队列来存储当前层的节点
//	index := 1                 // 从数组的第二个元素开始处理
//
//	for index < len(arr) {
//		// 取队列头节点作为当前父节点
//		current := queue[0]
//		queue = queue[1:]
//
//		// 处理左子节点
//		if arr[index] != nil {
//			leftNode := &TreeNode{Val: arr[index].(int)}
//			current.Left = leftNode
//			queue = append(queue, leftNode)
//		}
//		index++
//
//		// 处理右子节点
//		if index < len(arr) && arr[index] != nil {
//			rightNode := &TreeNode{Val: arr[index].(int)}
//			current.Right = rightNode
//			queue = append(queue, rightNode)
//		}
//		index++
//	}
//
//	return root
//}
//
//func printTree(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	fmt.Printf("%d ", root.Val)
//	printTree(root.Left)
//	printTree(root.Right)
//}
//
//func main() {
//	arr := []interface{}{1, 2, 2, 3, 4, 4, 3}
//	root := BuildTreeFromArray(arr)
//	//fmt.Println("Build binary-tree:")
//
//	result := findBottomLeftValue(root)
//
//	fmt.Println(result)
//}
//
//func findBottomLeftValue(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	maxHeight := findHeight(root)
//
//	queue := []*TreeNode{root}
//	depth := 1
//	size := 1
//	var result int
//
//	for len(queue) > 0 {
//		//fmt.Println(queue[0].Val)
//		if depth == maxHeight {
//			result = queue[0].Val
//			break
//		}
//		var subSize int
//		for size > 0 {
//			pop := queue[0]
//			queue = queue[1:]
//			size--
//
//			if pop.Left != nil {
//				queue = append(queue, pop.Left)
//				subSize++
//			}
//			if pop.Right != nil {
//				queue = append(queue, pop.Right)
//				subSize++
//			}
//		}
//		size +=subSize
//		depth++
//	}
//
//	return result
//}
//
//func findHeight(node *TreeNode) int {
//	if node == nil {
//		return 0
//	}
//
//	return max(findHeight(node.Left), findHeight(node.Right)) + 1
//}
