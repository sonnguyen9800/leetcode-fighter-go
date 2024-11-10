package main

func IsAnagram(first_str string, second_str string) bool {
	if len(first_str) != len(second_str) {
		return false
	}

	runeMap_first, runMap_second := make(map[rune]int), make(map[rune]int)

	for i := 0; i < len(first_str); i++ {
		runeMap_first[rune(first_str[i])]++
		runMap_second[rune(first_str[i])]++
	}

	for currentRune, value := range runeMap_first {
		value_2, exist := runMap_second[currentRune]
		if !exist || value != value_2 {
			return false
		}
	}
	return true
}
