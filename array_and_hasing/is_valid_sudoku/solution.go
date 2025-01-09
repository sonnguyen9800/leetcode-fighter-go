package is_valid_sudoku

import "strconv"

func parsed(input []string) []int {
	array := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		array[i], _ = strconv.Atoi(input[i])
	}
	return array
}
func is_valid_sudoku(input [][]string) bool {

	parsed_input := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		parsed_input[i] = parsed(input[i])
	}
	hashmap := make(map[int]bool)
	// Check all row
	for i := 0; i < len(parsed_input); i++ {
		hashmap = make(map[int]bool)

		for j := 0; j < len(parsed_input[i]); j++ {
			if parsed_input[i][j] == 0 {
				continue
			}
			if hashmap[parsed_input[i][j]] == true {
				return false
			}
			hashmap[parsed_input[i][j]] = true
		}
	}

	//Check all column
	for i := 0; i < len(parsed_input); i++ {
		hashmap = make(map[int]bool)

		for j := 0; j < len(parsed_input[i]); j++ {
			if parsed_input[j][i] == 0 {
				continue
			}
			if hashmap[parsed_input[j][i]] == true {
				return false
			}
			hashmap[parsed_input[j][i]] = true
		}
	}
	// Check each square
	count := 0
	for i, j := 0, 0; i < len(parsed_input) && j < len(parsed_input); {
		hashmap = make(map[int]bool)

		for sqrI := 0; sqrI < 3; sqrI++ {
			for sqrJ := 0; sqrJ < 3; sqrJ++ {
				if parsed_input[i+sqrI][sqrJ+j] == 0 {
					continue
				}
				if hashmap[parsed_input[i+sqrI][sqrJ+j]] == true {
					return false
				} else {
					hashmap[parsed_input[i+sqrI][sqrJ+j]] = true
				}
			}
		}
		i += 3
		count++
		if count == 3 {
			count = 0
			j += 3
			i = 0
		}

		//j += 3
	}

	return true
}

func is_valid_sudoku_optimal(input [][]string) bool {

	parsed_input := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		parsed_input[i] = parsed(input[i])
	}
	hashmap_row := make(map[int]bool, 9)
	hashmap_col := make(map[int]bool, 9)
	hashmap_sqr := make(map[int]bool, 9)
	// Check all row
	for i := 0; i < len(parsed_input); i++ {

		for j := 0; j < len(parsed_input[i]); j++ {

			if parsed_input[i][j] != 0 && hashmap_row[parsed_input[i][j]] == true {
				return false
			}
			if parsed_input[j][i] != 0 && hashmap_col[parsed_input[j][i]] == true {
				return false
			}

			sqr_index := int(i/3)*3 + int(j/3)
			if parsed_input[i][j] != 0 && hashmap_sqr[sqr_index] == true {
				return false
			}
			hashmap_sqr[sqr_index] = true
			hashmap_row[parsed_input[i][j]] = true
			hashmap_col[parsed_input[j][j]] = true
		}
	}

	return true
}
