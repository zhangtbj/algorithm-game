package main

import (
	"fmt"
)

var board [][]string

var nums = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

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
	backtracking(board)
}

func backtracking(board [][]string) bool {
	for i := 0; i < 81; i++ {
		row := i / 9
		col := i % 9
		//fmt.Printf("row: %d, col: %d\n", row, col)

		if board[row][col] != "." {
			continue
		}

		for j := range nums {
			if isvalid(board, row, col, nums[j]) {
				board[row][col] = nums[j]
				fmt.Println(board)
				if backtracking(board) == true {
					return true
				}
				board[row][col] = "."
			}
		}
		//9个数都尝试过了都不行，就return false
		return false
	}
	return true
}

func isvalid(board [][]string, row int, col int, k string) bool {

	for i := 0; i < 9; i++ {
		if board[row][i] == k {
			return false
		}
		if board[i][col] == k {
			return false
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
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
