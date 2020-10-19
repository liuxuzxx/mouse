package leetcode

import "fmt"

//37. Sudoku Solver
//Hard
//
//2069
//
//93
//
//Add to List
//
//Share
//Write a program to solve a Sudoku puzzle by filling the empty cells.
//
//A sudoku solution must satisfy all of the following rules:
//
//Each of the digits 1-9 must occur exactly once in each row.
//Each of the digits 1-9 must occur exactly once in each column.
//Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
//Empty cells are indicated by the character '.'.
//
//
//A sudoku puzzle...
//
//
//...and its solution numbers marked in red.
//
//Note:
//
//The given board contain only digits 1-9 and the character '.'.
//You may assume that the given Sudoku puzzle will have a single unique solution.
//The given board size is always 9x9.

func SudokuSolver() {
	board := [][]byte{
		{5, 3, byte('.'), byte('.'), 7, byte('.'), byte('.'), byte('.'), byte('.')},
		{6, byte('.'), byte('.'), 1, 9, 5, byte('.'), byte('.'), byte('.')},
		{byte('.'), 9, 8, byte('.'), byte('.'), byte('.'), byte('.'), 6, byte('.')},
		{8, byte('.'), byte('.'), byte('.'), 6, byte('.'), byte('.'), byte('.'), 3},
		{4, byte('.'), byte('.'), 8, byte('.'), 3, byte('.'), byte('.'), 1},
		{7, byte('.'), byte('.'), byte('.'), 2, byte('.'), byte('.'), byte('.'), 6},
		{byte('.'), 6, byte('.'), byte('.'), byte('.'), byte('.'), 2, 8, byte('.')},
		{byte('.'), byte('.'), byte('.'), 4, 1, 9, byte('.'), byte('.'), 5},
		{byte('.'), byte('.'), byte('.'), byte('.'), 8, byte('.'), byte('.'), 7, 9},
	}

	doSudokuSolver(&board)
	fmt.Printf("结果是")
}

var (
	rowMap    = make([]map[byte]interface{}, 9)
	columnMap = make([]map[byte]interface{}, 9)
	blockMap  = make([]map[byte]interface{}, 9)
)

func doSudokuSolver(boardPointer *[][]byte) {
	board := *boardPointer
	for index := 0; index < 9; index = index + 1 {
		rowMap[index] = make(map[byte]interface{}, 9)
		columnMap[index] = make(map[byte]interface{}, 9)
		blockMap[index] = make(map[byte]interface{}, 9)
	}

	for x := 0; x < 9; x = x + 1 {
		for y := 0; y < 9; y = y + 1 {
			element := board[x][y]
			if element != '.' {
				columnMap[x][board[x][y]] = true
				rowMap[y][board[x][y]] = true
				blockMap[x/3+y/3*3][board[x][y]] = true
			}
		}
	}

	resolve(boardPointer)

}

func resolve(boardPointer *[][]byte) bool {
	board := *boardPointer
	for x := 0; x < 9; x = x + 1 {
		for y := 0; y < 9; y = y + 1 {
			element := board[x][y]
			if element == '.' {
				for number := 1; number <= 9; number = number + 1 {
					_, rowOk := rowMap[y][byte(number)]
					_, columnOk := columnMap[x][byte(number)]
					_, blockOk := blockMap[x/3+y/3*3][byte(number)]
					if rowOk || columnOk || blockOk {
						continue
					}
					board[x][y] = byte(number)
					columnMap[x][board[x][y]] = true
					rowMap[y][board[x][y]] = true
					blockMap[x/3+y/3*3][board[x][y]] = true
					if resolve(boardPointer) {
						return true
					} else {
						tempElement := board[x][y]
						delete(rowMap[y], tempElement)
						delete(columnMap[x], tempElement)
						delete(blockMap[x/3+y/3*3], tempElement)
						board[x][y] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}
