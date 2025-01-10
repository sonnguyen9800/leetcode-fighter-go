package is_valid_sudoku

func isValidSudoku_neetcode(board [][]string) bool {
	rows := make([]int, 9)
	cols := make([]int, 9)
	squares := make([]int, 9)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == "." {
				continue
			}
			runeSlice := []rune(board[r][c])

			val := runeSlice[0] - '1'
			bit := 1 << val
			squareIdx := (r/3)*3 + c/3

			if rows[r]&bit != 0 || cols[c]&bit != 0 ||
				squares[squareIdx]&bit != 0 {
				return false
			}

			rows[r] |= bit
			cols[c] |= bit
			squares[squareIdx] |= bit
		}
	}

	return true
}

func isValidSudoku_Neetcode2(board [][]string) bool {
	rows := make([]map[string]bool, 9)
	cols := make([]map[string]bool, 9)
	squares := make([]map[string]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[string]bool)
		cols[i] = make(map[string]bool)
		squares[i] = make(map[string]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == "." {
				continue
			}
			val := board[r][c]
			squareIdx := (r/3)*3 + c/3

			if rows[r][val] || cols[c][val] ||
				squares[squareIdx][val] {
				return false
			}

			rows[r][val] = true
			cols[c][val] = true
			squares[squareIdx][val] = true
		}
	}

	return true
}
