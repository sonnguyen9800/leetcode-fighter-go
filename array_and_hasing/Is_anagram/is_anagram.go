package is_anagram

import "strings"

func IsAnagram(inputString InputString) bool {
	first_str := strings.ToLower(inputString.S1)
	second_str := strings.ToLower(inputString.S2)
	if len(first_str) != len(second_str) {
		return false
	}

	runeMap_first, runMap_second := make(map[rune]int), make(map[rune]int)

	for i := 0; i < len(first_str); i++ {
		runeMap_first[rune(first_str[i])]++
		runMap_second[rune(second_str[i])]++
	}

	for currentRune, value := range runeMap_first {
		value_2, exist := runMap_second[currentRune]
		if !exist || value != value_2 {
			return false
		}
	}
	return true
}

func IsAnagramOptimal(inputString InputString) bool {
	first_str := strings.ToLower(inputString.S1)
	second_str := strings.ToLower(inputString.S2)
	if len(first_str) != len(second_str) {
		return false
	}
	if len(first_str) != len(second_str) {
		return false
	}

	runeMap := make(map[rune]int)
	for i := 0; i < len(first_str); i++ {
		runeMap[rune(first_str[i])] += 1
		runeMap[rune(second_str[i])] -= 1
	}
	for _, value := range runeMap {
		if value != 0 {
			return false
		}
	}
	return true
}
