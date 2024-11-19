package main

func TwoSumBruteForce(input TwoIntergersInput) []int {
	num_array := input.ArrayInt
	target := input.Target

	dest1 := make([]int, len(num_array))
	copy(dest1, num_array)

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

func _TwoSum(num_array []int, target int) []int {

	dest1 := make([]int, len(num_array))
	copy(dest1, num_array)

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
