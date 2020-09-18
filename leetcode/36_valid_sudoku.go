package leetcode

import "fmt"

//36. Valid Sudoku
//Medium
//
//1881
//
//477
//
//Add to List
//
//Share
//Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
//Each row must contain the digits 1-9 without repetition.
//Each column must contain the digits 1-9 without repetition.
//Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//
//A partially filled sudoku which is valid.
//
//The Sudoku board could be partially filled, where empty cells are filled with the character '.'.
//
//Example 1:
//
//Input:
//[
//["5","3",byte('.'),byte('.'),"7",byte('.'),byte('.'),byte('.'),byte('.')},
//["6",byte('.'),byte('.'),"1","9","5",byte('.'),byte('.'),byte('.')},
//[byte('.'),"9","8",byte('.'),byte('.'),byte('.'),byte('.'),"6",byte('.')},
//["8",byte('.'),byte('.'),byte('.'),"6",byte('.'),byte('.'),byte('.'),"3"},
//["4",byte('.'),byte('.'),"8",byte('.'),"3",byte('.'),byte('.'),"1"},
//["7",byte('.'),byte('.'),byte('.'),"2",byte('.'),byte('.'),byte('.'),"6"},
//[byte('.'),"6",byte('.'),byte('.'),byte('.'),byte('.'),"2","8",byte('.')},
//[byte('.'),byte('.'),byte('.'),"4","1","9",byte('.'),byte('.'),"5"},
//[byte('.'),byte('.'),byte('.'),byte('.'),"8",byte('.'),byte('.'),"7","9"]
//]
//Output: true
//Example 2:
//
//Input:
//[
//["8","3",byte('.'),byte('.'),"7",byte('.'),byte('.'),byte('.'),byte('.')},
//["6",byte('.'),byte('.'),"1","9","5",byte('.'),byte('.'),byte('.')},
//[byte('.'),"9","8",byte('.'),byte('.'),byte('.'),byte('.'),"6",byte('.')},
//["8",byte('.'),byte('.'),byte('.'),"6",byte('.'),byte('.'),byte('.'),"3"},
//["4",byte('.'),byte('.'),"8",byte('.'),"3",byte('.'),byte('.'),"1"},
//["7",byte('.'),byte('.'),byte('.'),"2",byte('.'),byte('.'),byte('.'),"6"},
//[byte('.'),"6",byte('.'),byte('.'),byte('.'),byte('.'),"2","8",byte('.')},
//[byte('.'),byte('.'),byte('.'),"4","1","9",byte('.'),byte('.'),"5"},
//[byte('.'),byte('.'),byte('.'),byte('.'),"8",byte('.'),byte('.'),"7","9"]
//]
//Output: false
//Explanation: Same as Example 1, except with the 5 in the top left corner being
//modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
//Note:
//
//A Sudoku board (partially filled) could be valid but is not necessarily solvable.
//Only the filled cells need to be validated according to the mentioned rules.
//The given board contain only digits 1-9 and the character '.'.
//The given board size is always 9x9.

func ValidSudoku() {
	board := [][]byte{
		{8, 3, byte('.'), byte('.'), 7, byte('.'), byte('.'), byte('.'), byte('.')},
		{6, byte('.'), byte('.'), 1, 9, 5, byte('.'), byte('.'), byte('.')},
		{byte('.'), 9, 8, byte('.'), byte('.'), byte('.'), byte('.'), 6, byte('.')},
		{8, byte('.'), byte('.'), byte('.'), 6, byte('.'), byte('.'), byte('.'), 3},
		{4, byte('.'), byte('.'), 8, byte('.'), 3, byte('.'), byte('.'), 1},
		{7, byte('.'), byte('.'), byte('.'), 2, byte('.'), byte('.'), byte('.'), 6},
		{byte('.'), 6, byte('.'), byte('.'), byte('.'), byte('.'), 2, 8, byte('.')},
		{byte('.'), byte('.'), byte('.'), 4, 1, 9, byte('.'), byte('.'), 5},
		{byte('.'), byte('.'), byte('.'), byte('.'), 8, byte('.'), byte('.'), 7, 9},
	}

	result := doValidSudoku(board)
	fmt.Printf("判断结果是:%s", result)
}

func doValidSudoku(board [][]byte) bool {
	for x := 0; x < 9; x = x + 1 {
		rowMap := make(map[byte]interface{}, 9)
		for y := 0; y < 9; y = y + 1 {
			element := board[x][y]
			if _, ok := rowMap[element]; element != '.' && ok {
				return false
			} else {
				rowMap[element] = true
			}
		}
	}

	for y := 0; y < 9; y = y + 1 {
		columnMap := make(map[byte]interface{}, 9)
		for x := 0; x < 9; x = x + 1 {
			element := board[x][y]
			if _, ok := columnMap[element]; element != '.' && ok {
				return false
			} else {
				columnMap[element] = true
			}
		}
	}

	for xStep := 1; xStep <= 3; xStep = xStep + 1 {
		for yStep := 1; yStep <= 3; yStep = yStep + 1 {
			blockMap := make(map[byte]interface{}, 9)
			for x := (xStep - 1) * 3; x < xStep*3; x = x + 1 {
				for y := (yStep - 1) * 3; y < yStep*3; y = y + 1 {
					element := board[x][y]
					if _, ok := blockMap[element]; element != '.' && ok {
						return false
					} else {
						blockMap[element] = true
					}
				}
			}
		}
	}

	return true
}
