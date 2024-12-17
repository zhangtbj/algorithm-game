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
//	result := isSymmetric(root)
//
//	fmt.Println(result)
//}
//
//func isSymmetric(root *TreeNode) bool {
//	var nodes []*TreeNode
//	var depth = 1
//	nodes = append(nodes, root)
//
//	if root == nil {
//		return true
//	}
//
//	var first = true
//	var pop *TreeNode
//	for len(nodes) > 0 {
//		var subDepth int
//		var depthArr []int
//		for depth > 0 {
//			pop = nodes[0]
//			nodes = nodes[1:]
//			if pop != nil {
//				depthArr = append(depthArr, pop.Val)
//			} else {
//				depthArr = append(depthArr, -200)
//			}
//
//			depth--
//
//			if pop != nil {
//				nodes = append(nodes, pop.Left)
//				nodes = append(nodes, pop.Right)
//				subDepth += 2
//			}
//		}
//
//		fmt.Println(depthArr)
//
//		if !first {
//			if len(depthArr)%2 != 0 {
//				return false
//			} else {
//				start, end := 0, len(depthArr)-1
//				for start < end {
//					if depthArr[start] != depthArr[end] {
//						return false
//					}
//					start++
//					end--
//				}
//			}
//		}
//		first = false
//		depth += subDepth
//	}
//
//	return true
//}
