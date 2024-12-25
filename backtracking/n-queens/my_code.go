package main

//import (
//	"fmt"
//)
//
//var result [][]string
//var queenMap [][]int
//var path []int
//
//func main() {
//	n := 4
//	solveNQueens(n)
//	fmt.Println(result)
//}
//
//func solveNQueens(n int) [][]string {
//	result = [][]string{}
//	queenMap = make([][]int, n)
//	for i := range queenMap {
//		queenMap[i] = make([]int, n)
//	}
//	path = []int{}
//	backtracking(n, queenMap, path, 0)
//	return result
//}
//
//func backtracking(n int, queenMap [][]int, path []int, row int) {
//	if len(path) == n {
//		tmp := make([]string, n)
//		for i := range path {
//			var pathStr string
//			for j := 0; j < n; j++ {
//				if path[i] != j {
//					pathStr += "."
//				} else {
//					pathStr += "Q"
//				}
//			}
//			tmp[i] = pathStr
//		}
//		result = append(result, tmp)
//		return
//	}
//
//	for i := 0; i < n; i++ {
//		if queenMap[row][i] > 0 {
//			continue
//		}
//		path = append(path, i)
//		queenMap = addQueue(queenMap, n, row, i, true)
//		backtracking(n, queenMap, path, row+1)
//		queenMap = addQueue(queenMap, n, row, i, false)
//		path = path[:len(path)-1]
//	}
//}
//
//func addQueue(queenMap [][]int, n int, row int, col int, add bool) [][]int {
//	for i := 0; i < n; i++ {
//		if add {
//			queenMap[row][col] = 1
//			if i != col {
//				queenMap[row][i]++
//			}
//
//			if i != row {
//				queenMap[i][col]++
//			}
//		} else {
//			queenMap[row][i]--
//			queenMap[i][col]--
//			queenMap[row][col] = 0
//		}
//	}
//
//	for i := 1; i < n; i++ {
//		if row+i == n || col+i == n {
//			break
//		}
//		if add {
//			queenMap[row+i][col+i]++
//		} else {
//			queenMap[row+i][col+i]--
//		}
//	}
//
//	for i := 1; i < n; i++ {
//
//	}
//
//	for i := 1; i < n; i++ {
//		if row+i == n || col-i < 0 {
//			break
//		}
//		if add {
//			queenMap[row+i][col-i]++
//		} else {
//			queenMap[row+i][col-i]--
//		}
//	}
//
//	for i := 1; i < n; i++ {
//		if row-i < 0 || col-i < 0 {
//			break
//		}
//		if add {
//			queenMap[row-i][col-i]++
//		} else {
//			queenMap[row-i][col-i]--
//		}
//	}
//
//	return queenMap
//}
