package anagram_groups

import is_anagram "leetcode-fighter-go/array_and_hasing/Is_anagram"

func groupAnagramsNeetCodeSolution(strs []string) [][]string {
	res := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		res[count] = append(res[count], s)
	}

	var result [][]string
	for _, group := range res {
		result = append(result, group)
	}
	return result
}

func CheckIfAnagramExist(string2 string, map_input map[string][]string) {
	exist_key := false
	for key, _ := range map_input {
		if (is_anagram.IsAnagram(is_anagram.InputString{
			S1: string2,
			S2: key,
		})) {
			map_input[key] = append(map_input[key], string2)
			exist_key = true
		}
	}
	if !exist_key {
		map_input[string2] = []string{string2}
	}

}

func GroupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		CheckIfAnagramExist(str, anagramMap)
	}

	data := [][]string{}

	for _, strs_value := range anagramMap {
		array_str := []string{}
		for _, str := range strs_value {
			array_str = append(array_str, str)
		}
		data = append(data, array_str)
	}

	return data
}
