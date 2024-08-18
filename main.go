package main

import (
	"errors"
	"fmt"
)

func findNextEmpty(board [][]byte) (byte, byte, error) {
	for j := byte(0); j < 9; j++ {
		for i := byte(0); i < 9; i++ {
			if board[j][i] == '.' {
				return j, i, nil
			}
		}
	}
	return byte(0), byte(0), errors.New("no empty cell to fill in")
}

func ValidRule(board [][]byte, j, i byte) bool {
	num := board[j][i]
	// Check if the number is valid in the row
	for k := byte(0); k < 9; k++ {
		if board[j][k] == num && k != i {
			return false
		}
	}

	// Check if the number is valid in the column
	for k := byte(0); k < 9; k++ {
		if board[k][i] == num && k != j {
			return false
		}
	}

	// Check if the number is valid in the 3x3 grid
	for k := byte(0); k < 3; k++ {
		for l := byte(0); l < 3; l++ {
			if board[j-j%3+k][i-i%3+l] == num && (j-j%3+k != j || i-i%3+l != i) {
				return false
			}
		}
	}
	return true
}

func backtracing(board [][]byte, j, i byte) bool {
	// fill in avabilable choices
	for k := byte(1); k <= 9; k++ {
		board[j][i] = 48 + k
		if ValidRule(board, j, i) {
			j, i, e := findNextEmpty(board)
			if e != nil {
				return true
			}
			status := backtracing(board, j, i)
			if status {
				return true
			}
		}
		board[j][i] = '.'
	}
	return false
}

func solveSudoku(board [][]byte) {
	j, i, e := findNextEmpty(board)
	if e != nil {
		return
	}
	// start backtracking
	backtracing(board, j, i)
}

func printBoard(board [][]byte) {
	for _, row := range board {
		for ix, col := range row {
			if ix == 8 {
				fmt.Printf("%c\n", col)
			} else {
				fmt.Printf("%c ", col)
			}
		}
	}
}

func main() {
	fmt.Println("")
	fmt.Println("Found Solution")
	fmt.Println("")
	board := [][]byte{
		{'6', '.', '.', '.', '.', '2', '5', '.', '.'},
		{'.', '1', '7', '5', '.', '.', '.', '.', '.'},
		{'4', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '7', '.', '.', '2', '3', '.', '6', '.'},
		{'.', '.', '.', '.', '1', '.', '3', '.', '.'},
		{'.', '.', '2', '.', '.', '5', '7', '.', '.'},
		{'.', '.', '.', '4', '.', '.', '.', '.', '.'},
		{'.', '9', '5', '.', '.', '.', '.', '3', '.'},
		{'1', '.', '8', '.', '.', '.', '9', '.', '.'},
	}
	solveSudoku(board)
	printBoard(board)
	fmt.Println("")
	fmt.Println("CREDIT : ANASS AJJA")
	fmt.Println("")
}
