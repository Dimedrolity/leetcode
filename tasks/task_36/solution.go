// Package task_36
// 36. Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/description/
package task_36

var emptyCell = byte('.')

// isValidSudoku
// Мое решение.
// Сложность по времени O(N), по памяти O(N), где N - кол-во ячеек доски.
// Идея в том, чтобы проверить строки на наличие дубликатов с помощью set'а, аналогично столбцы и квадраты.
func isValidSudoku(board [][]byte) bool {
	isValid := rowsValid(board)
	if !isValid {
		return false
	}

	isValid = columnsValid(board)
	if !isValid {
		return false
	}

	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board); j += 3 {
			isValid = boardPartValid(board, i, i+3, j, j+3)
			if !isValid {
				return false
			}

		}
	}

	return true
}

// boardPartValid возвращает true, если часть доски по индексам строк и индексам столбцов не содержит дубликат.
func boardPartValid(board [][]byte, startI int, endI int, startJ int, endJ int) bool {
	set := make(map[byte]bool, 9)
	for i := startI; i < endI; i++ {
		for j := startJ; j < endJ; j++ {
			currentCell := board[i][j]

			if currentCell == emptyCell {
				continue
			}

			_, duplicates := set[currentCell]
			if duplicates {
				return false
			} else {
				set[currentCell] = true
			}
		}
	}
	return true
}

// columnsValid возвращает true, если в столбцы не содержат дубликат.
func columnsValid(board [][]byte) bool {
	for j := 0; j < len(board); j++ {
		isValid := boardPartValid(board, 0, 9, j, j+1)
		if !isValid {
			return false
		}
	}

	return true
}

// rowsValid возвращает true, если в строки не содержат дубликат.
func rowsValid(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		isValid := boardPartValid(board, i, i+1, 0, 9)
		if !isValid {
			return false
		}
	}

	return true
}
