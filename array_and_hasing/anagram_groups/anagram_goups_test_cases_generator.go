package anagram_groups

import (
	"math/rand"
	"sort"
)

// generateRandomString creates a random string of a given length.
func generateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// shuffleString shuffles the characters in a string.
func shuffleString(s string) string {
	runes := []rune(s)
	rand.Shuffle(len(runes), func(i, j int) { runes[i], runes[j] = runes[j], runes[i] })
	return string(runes)
}

// generateAnagramGroup generates a list of anagrams from a base word.
func generateAnagramGroup(base string, count int) []string {
	anagrams := make([]string, count)
	for i := range anagrams {
		anagrams[i] = shuffleString(base)
	}
	return anagrams
}
func fuzzyTestCaseGeneratorAdv(
	min_group_num int, max_group_num int,
	group_size_min int, group_size_max int,
	random_noise_string int) ([]string, [][]string) {
	var allStrings []string
	var expectedOutput [][]string

	numGroups := rand.Intn(max_group_num) + min_group_num
	groupSize := rand.Intn(group_size_max) + group_size_min
	// Generate anagram groups
	for i := 0; i < numGroups; i++ {
		base := generateRandomString(rand.Intn(6) + 3) // Base word length between 3 and 8
		group := generateAnagramGroup(base, groupSize)
		allStrings = append(allStrings, group...)
		expectedOutput = append(expectedOutput, group)
	}
	random_noise := rand.Intn(random_noise_string)
	// Generate random non-anagram strings
	for i := 0; i < random_noise; i++ {
		allStrings = append(allStrings, generateRandomString(rand.Intn(6)+3))
	}

	// Shuffle all input strings
	rand.Shuffle(len(allStrings), func(i, j int) { allStrings[i], allStrings[j] = allStrings[j], allStrings[i] })

	return allStrings, expectedOutput
}

// fuzzyTestCaseGenerator generates test cases for the "Group Anagrams" problem.
func fuzzyTestCaseGenerator(numGroups, groupSize, randomStrings int) ([]string, [][]string) {
	var allStrings []string
	var expectedOutput [][]string

	// Generate anagram groups
	for i := 0; i < numGroups; i++ {
		base := generateRandomString(rand.Intn(6) + 3) // Base word length between 3 and 8
		group := generateAnagramGroup(base, groupSize)
		allStrings = append(allStrings, group...)
		expectedOutput = append(expectedOutput, group)
	}

	// Generate random non-anagram strings
	for i := 0; i < randomStrings; i++ {
		allStrings = append(allStrings, generateRandomString(rand.Intn(6)+3))
	}

	// Shuffle all input strings
	rand.Shuffle(len(allStrings), func(i, j int) { allStrings[i], allStrings[j] = allStrings[j], allStrings[i] })

	return allStrings, expectedOutput
}

// normalizeOutput groups the strings into anagrams for validation purposes.
func normalizeOutput(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		sorted := sortString(str)
		anagramMap[sorted] = append(anagramMap[sorted], str)
	}

	var result [][]string
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

// sortString returns the sorted version of a string (used as a key).
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}
