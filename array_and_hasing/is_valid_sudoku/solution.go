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
	for i, j := 0, 0; i < len(parsed_input) && j < len(parsed_input); i += 3 {
		hashmap = make(map[int]bool)

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if parsed_input[i][j] == 0 {
					continue
				}
				if hashmap[parsed_input[i][j]] == true {
					return false
				} else {
					hashmap[parsed_input[i][j]] = true
				}
			}
		}
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
