package main

// import (
//
//	"fmt"
//
// )
//
//	type TreeNode struct {
//		Val   int
//		Left  *TreeNode
//		Right *TreeNode
//	}
//
//	func BuildTreeFromArray(arr []interface{}) *TreeNode {
//		if len(arr) == 0 {
//			return nil
//		}
//
//		// 将根节点初始化
//		root := &TreeNode{Val: arr[0].(int)}
//		queue := []*TreeNode{root} // 用队列来存储当前层的节点
//		index := 1                 // 从数组的第二个元素开始处理
//
//		for index < len(arr) {
//			// 取队列头节点作为当前父节点
//			current := queue[0]
//			queue = queue[1:]
//
//			// 处理左子节点
//			if arr[index] != nil {
//				leftNode := &TreeNode{Val: arr[index].(int)}
//				current.Left = leftNode
//				queue = append(queue, leftNode)
//			}
//			index++
//
//			// 处理右子节点
//			if index < len(arr) && arr[index] != nil {
//				rightNode := &TreeNode{Val: arr[index].(int)}
//				current.Right = rightNode
//				queue = append(queue, rightNode)
//			}
//			index++
//		}
//
//		return root
//	}
//
//	func printTree(root *TreeNode) {
//		if root == nil {
//			return
//		}
//		fmt.Printf("%d ", root.Val)
//		printTree(root.Left)
//		printTree(root.Right)
//	}
//
//	func main() {
//		arr := []interface{}{1,null,2,2}
//		root := BuildTreeFromArray(arr)
//		//fmt.Println("Build binary-tree:")
//
//		result := findMode(root)
//
//		fmt.Println(result)
//	}
//func findMode(root *TreeNode) []int {
//	var nodes []int
//	var result []int
//	inorder(root, &nodes)
//
//	var maxNodeCount = 0
//	var nodeMap = make(map[int]int)
//	for i := range nodes {
//		nodeMap[nodes[i]]++
//	}
//
//	for _, v := range nodeMap {
//		maxNodeCount = max(v, maxNodeCount)
//	}
//
//	for k, v := range nodeMap {
//		if v == maxNodeCount {
//			result = append(result, k)
//		}
//	}
//
//	return result
//}
//
//func inorder(root *TreeNode, arr *[]int) {
//	if root == nil {
//		return
//	}
//
//	inorder(root.Left, arr)
//	*arr = append(*arr, root.Val)
//	inorder(root.Right, arr)
//}