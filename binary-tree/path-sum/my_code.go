package main

//import (
//	"fmt"
//	"os"
//)
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
//	arr := []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}
//	root := BuildTreeFromArray(arr)
//	//fmt.Println("Build binary-tree:")
//
//	target := 22
//	result := hasPathSum(root, target)
//
//	fmt.Println(result)
//}
//
//func hasPathSum(root *TreeNode, targetSum int) bool {
//	fmt.Printf("targetSum: %d\n", targetSum)
//	if root == nil {
//		return false
//	}
//
//	if root.Left == nil && root.Right == nil {
//		if targetSum-root.Val == 0 {
//			return true
//			os.Exit(0)
//		} else {
//			return false
//		}
//	}
//
//	var leftHas bool
//	if root.Left != nil {
//		targetSum = targetSum - root.Val
//		leftHas = hasPathSum(root.Left, targetSum)
//		targetSum = targetSum + root.Val
//	}
//
//	var rightHas bool
//	if root.Right != nil {
//		targetSum = targetSum - root.Val
//		rightHas = hasPathSum(root.Right, targetSum)
//		targetSum = targetSum + root.Val
//	}
//
//	return leftHas || rightHas
//}
