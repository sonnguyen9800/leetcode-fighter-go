package main

func IsAnagram(first_str string, second_str string) bool {
	if len(first_str) != len(second_str) {
		return false
	}

	runeMap := make(map[rune]int)

	var first_str_array = []rune(first_str)
	for _, rune := range first_str_array {
		value, exists := runeMap[rune]
		if !exists {
			value = 1
			runeMap[rune] = value
		} else {
			value += 1
			runeMap[rune] = value
		}
	}
	second_str_array := []rune(second_str)
	runmap_second_str := make(map[rune]int)

	for _, rune := range second_str_array {
		_, exists := runeMap[rune]
		if !exists {
			return false
		}
		runmap_second_str[rune] += 1

	}
	for rune, value := range runeMap {
		value_2, exist := runmap_second_str[rune]
		if !exist || value != value_2 {
			return false
		}
	}
	return true
}
