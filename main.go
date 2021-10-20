package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func initBoard() [][]int {

	board := [][]int{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}
	return board
}

func checkValid(row, col, val int, board [][]int) bool {

	var isValid bool
	isValid = checkRowValid(row, val, board)
	isValid = isValid && checkColValid(col, val, board)
	isValid = isValid && checkBoxValid(row, col, val, board)

	return isValid
}

func checkRowValid(row, val int, board [][]int) bool {
	for _, v := range board[row] {
		if val == v {
			return false
		}
	}

	return true
}

func checkColValid(col, val int, board [][]int) bool {

	var column []int

	for _, row := range board {
		if row[col] != 0 {
			column = append(column, row[col])
		}
	}

	for _, v := range column {
		if val == v {
			return false
		}
	}

	return true
}

func checkBoxValid(row, col, val int, board [][]int) bool {
	rowBox := row / 3
	colBox := col / 3

	rowIdx := []int{0 + 3*rowBox, 1 + 3*rowBox, 2 + 3*rowBox}
	colIdx := []int{0 + 3*colBox, 1 + 3*colBox, 2 + 3*colBox}

	for _, colV := range colIdx {
		for _, rowV := range rowIdx {
			if val == board[rowV][colV] {
				return false
			}
		}
	}

	return true
}

func modifyBoard(row, col, val int, board [][]int) [][]int {

	newBoard := deepCopy(board)
	newBoard[row][col] = val
	return newBoard
}

// helper function to make a deep copy of 2D slice
func deepCopy(a [][]int) [][]int {

	b := make([][]int, len(a))

	for i := range a {
		b[i] = make([]int, len(a[i]))
		copy(b[i], a[i])
	}

	return b
}

// output: row, col, isFinished
// isFinished is true if the board is complete
func findNextUnknownPos(board [][]int) (int, int, bool) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == 0 {
				return row, col, false
			}
		}
	}
	return -1, -1, true
}

func getPosValues(board [][]int, row, col int) *stack.Stack {
	s := stack.New()

	for i := 1; i < 10; i++ {
		if checkValid(row, col, i, board) {
			s.Push(i)
		}
	}
	return s
}

func solveBoardRecursive(board [][]int) ([][]int, bool) {
	row, col, done := findNextUnknownPos(board)

	// check if board is already complete
	if done {
		return board, true
	}

	posVals := getPosValues(board, row, col)

	//check if board cannot be solved
	if posVals.Len() == 0 {
		return board, false
	}

	for posVals.Len() > 0 {
		tryVal := posVals.Pop().(int)
		newBoard := modifyBoard(row, col, tryVal, board)
		solvedBoard, solved := solveBoardRecursive(newBoard)
		if solved {
			return solvedBoard, solved
		}
	}

	return board, false
}

func main() {

	board := initBoard()

	solvedB, solved := solveBoardRecursive(board)

	if solved {
		for _, row := range solvedB {
			fmt.Println(row)
		}
	} else {
		fmt.Println("Could not solve")
	}
}
