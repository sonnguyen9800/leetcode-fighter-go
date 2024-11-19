package main

func TwoSumBruteForce(input TwoIntergersInput) []int {
	num_array := input.ArrayInt
	target := input.Target

	for i := 0; i < len(num_array); i++ {
		for j := i + 0; j < len(num_array); j++ {
			if i == j {
				continue
			}
			if num_array[i]+num_array[j] == target {
				return []int{i, j}
			}

		}
	}

	return []int{0, 0}
}

func TwoSumOnePass(input TwoIntergersInput) []int {
	dict_input := make(map[int]int)
	for i := 0; i < len(input.ArrayInt); i++ {
		dict_input[input.ArrayInt[i]] = i
	}
	for i := 0; i < len(input.ArrayInt); i++ {
		value, exists := dict_input[input.Target-input.ArrayInt[i]]

		if value != i && exists {
			return []int{value, i}
		}
	}

	return []int{0, 0}
}
func TwoSumOnePassVer2(input TwoIntergersInput) []int {
	dict_input := make(map[int]int, len(input.ArrayInt))
	for i := 0; i < len(input.ArrayInt); i++ {
		dict_input[input.ArrayInt[i]] = i
	}
	for i := 0; i < len(input.ArrayInt); i++ {
		value, exists := dict_input[input.Target-input.ArrayInt[i]]

		if value != i && exists {
			return []int{value, i}
		}
	}

	return []int{0, 0}
}
func TwoSumOnePassVer3(input TwoIntergersInput) []int {
	dict_input := make(map[int]int, len(input.ArrayInt))

	for i := 0; i < len(input.ArrayInt); i++ {
		value, exists := dict_input[input.Target-input.ArrayInt[i]]

		if value != i && exists {
			return []int{value, i}
		}
		dict_input[input.ArrayInt[i]] = i
	}

	return nil
}
func TwoSumOnePassVer4_Backup2(input TwoIntergersInput) []int {
	dict_input := make(map[int]int)
	my_array := input.ArrayInt
	target := input.Target

	for i, num := range my_array {
		diff := target - num
		if value, exists := dict_input[diff]; exists {
			return []int{value, i}
		}
		dict_input[num] = i
	}

	return []int{0, 0}
}
func TwoSumOnePassVer4(input TwoIntergersInput) []int {
	dict_input := make(map[int]int)
	my_array := input.ArrayInt
	target := input.Target
	array_len := len(input.ArrayInt)

	for i := 0; i < array_len; i++ {
		array_value := my_array[i]
		value, exists := dict_input[target-array_value]

		if exists {
			return []int{value, i}
		}
		dict_input[array_value] = i
	}
	return []int{0, 0}
}
func TwoSumOnePassVer4_BACKUP(input TwoIntergersInput) []int {
	dict_input := make(map[int]int)
	array_len := len(input.ArrayInt)
	my_array := input.ArrayInt
	target := input.Target

	for i := 0; i < array_len; i++ {
		array_value := my_array[i]
		value, exists := dict_input[target-array_value]

		if exists {
			return []int{value, i}
		}
		dict_input[array_value] = i
	}

	for i, num := range my_array {
		value, exists := dict_input[target-num]

		if exists {
			return []int{value, i}
		}
		dict_input[num] = i
	}

	return []int{0, 0}
}
func TwoSumOnePassNeetcode(input TwoIntergersInput) []int {
	nums := input.ArrayInt

	target := input.Target
	prevMap := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if j, found := prevMap[diff]; found {
			return []int{j, i}
		}
		prevMap[n] = i
	}
	return []int{0, 0}
}
