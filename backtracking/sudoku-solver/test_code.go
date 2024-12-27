package main

import (
	"fmt"
)

var board [][]string

func main() {
	board = [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	solveSudoku(board)
	fmt.Println(board)
}

func solveSudoku(board [][]string) {
	backtracking(board, 0)
}

func backtracking(board [][]string, startIndex int) bool {
	for i := startIndex; i < 81; i++ {
		row := i / 9
		col := i % 9
		fmt.Printf("i: %d, row: %d, col: %d\n", i, row, col)

		if board[row][col] != "." {
			continue
		}

		nums := avaiableNums(board, row, col)
		//fmt.Println(nums)

		for j := 0; j < len(nums); j++ {
			board[row][col] = nums[j]
			fmt.Printf("i: %d\n", i)
			fmt.Println(board)
			if backtracking(board, i+1) {
				return true
			}

			board[row][col] = "."
		}
		return false
	}
	return true
}

func avaiableNums(board [][]string, row int, col int) []string {
	var nums []string
	used := make(map[string]bool, 9)
	used["1"] = false
	used["2"] = false
	used["3"] = false
	used["4"] = false
	used["5"] = false
	used["6"] = false
	used["7"] = false
	used["8"] = false
	used["9"] = false

	for i := 0; i < 9; i++ {
		if board[row][i] != "." {
			used[board[row][i]] = true
		}
		if board[i][col] != "." {
			used[board[i][col]] = true
		}
	}

	startR, endR, startC, endC := 0, 0, 0, 0
	if row < 3 {
		startR, endR = 0, 2
	} else if row < 6 {
		startR, endR = 3, 5
	} else {
		startR, endR = 6, 8
	}

	if col < 3 {
		startC, endC = 0, 2
	} else if col < 6 {
		startC, endC = 3, 5
	} else {
		startC, endC = 6, 8
	}

	for i := startR; i <= endR; i++ {
		for j := startC; j <= endC; j++ {
			if board[i][j] != "." {
				used[board[i][j]] = true
			}
		}
	}
	//fmt.Println(used)
	for k, v := range used {
		if !v {
			nums = append(nums, k)
		}
	}
	return nums
}
