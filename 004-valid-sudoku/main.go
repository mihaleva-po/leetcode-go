// https://leetcode.com/problems/valid-sudoku/description/
// Сложность: Medium

package main

func isValidSudoku(board [][]byte) bool {

	// строки
	for _, v := range board {

		stroka := make(map[byte]struct{}, 9)
		for _, item := range v {

			if _, ok := stroka[item]; ok {
				return false
			}

			if rune(item) != '.' {
				stroka[item] = struct{}{}
			}

		}
	}

	// столбцы
	for i := 0; i < 9; i++ {
		column := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if _, ok := column[board[j][i]]; ok {
				return false
			}

			if rune(board[j][i]) != '.' {
				column[board[j][i]] = struct{}{}
			}
		}

	}

	for k := range 3 {
		for n := range 3 {
			kvadrat := make(map[byte]struct{})
			for i := k * 3; i < 3*k+3; i++ {
				for j := n * 3; j < 3*n+3; j++ {
					if _, ok := kvadrat[board[i][j]]; ok {
						return false
					}

					if rune(board[i][j]) != '.' {
						kvadrat[board[i][j]] = struct{}{}
					}
				}
			}
		}

	}

	return true

}
