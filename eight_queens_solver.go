package sprint

import (
	"strconv"
	"strings"
)

func EightQueensSolver() string {
	var solutions []string
	board := make([]int, 8)
	solveQueens(board, 0, &solutions)
	return strings.Join(solutions, "\n")
}

func solveQueens(board []int, row int, solutions *[]string) {
	if row == 8 {
		*solutions = append(*solutions, formatSolution(board))
		return
	}

	for col := 0; col < 8; col++ {
		if isValid(board, row, col) {
			board[row] = col
			solveQueens(board, row+1, solutions)
		}
	}
}

func isValid(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || abs(row-i) == abs(col-board[i]) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func formatSolution(board []int) string {
	var solution strings.Builder
	for _, col := range board {
		solution.WriteString(strconv.Itoa(col + 1))
	}
	return solution.String()
}
