package main

import (
	"fmt"
	"leetcode-fighter-go/goland-slib/colors" // Replace with your actual module name
)

// Test runner function
func runTests() {
	testCases := []struct {
		s1, s2   string
		expected bool
	}{
		// Basic anagrams
		{"listen", "silent", true},
		{"anagram", "nagaram", true},
		{"rat", "tar", true},
		{"debitcard", "badcredit", true},
		{"evil", "vile", true},

		// Not anagrams
		{"hello", "world", false},
		{"cat", "rat", false},
		{"good", "dog", false},
		{"night", "thingy", false},
		{"listen", "listing", false},

		// Edge cases
		{"", "", true},             // Both strings empty
		{"a", "a", true},           // Single character, same
		{"a", "b", false},          // Single character, different
		{"ab", "a", false},         // Different lengths
		{"aabbcc", "abcabc", true}, // Repeating characters

		// Mixed case (assuming case-sensitive)
		{"Hello", "hello", false},
		{"Tea", "Eat", false},

		// Unicode characters
		{"你好", "好你", true},
		{"你好", "你你", false},

		// Long strings
		{"thequickbrownfoxjumpsoverthelazydog", "dogthelazyoverjumpsfoxquickthebrown", true},
		{string(make([]byte, 10000, 10000)), string(make([]byte, 10000, 10000)), true}, // Large identical strings
	}

	successCount := 0
	failureCount := 0

	for i, testCase := range testCases {
		result := IsAnagram(testCase.s1, testCase.s2)
		if result != testCase.expected {
			failureCount++
			fmt.Printf("%sTest Case %d Failed: Expected %v, Got %v\n", colors.Red(""), i+1, testCase.expected, result)
			fmt.Printf("Strings: s1=%q, s2=%q %s\n", testCase.s1, testCase.s2, colors.Reset(""))
		} else {
			successCount++
			fmt.Printf("%sTest Case %d Passed!%s\n", colors.Green(""), i+1, colors.Reset(""))
		}
	}

	// Summary
	totalTests := len(testCases)
	fmt.Printf("\n%sSummary:%s %s%d/%d tests passed%s (%.2f%% success rate)\n", colors.Green(""), colors.Reset(""), colors.Bold(""), successCount, totalTests, colors.Reset(""), float64(successCount)/float64(totalTests)*100)
	if failureCount > 0 {
		fmt.Printf("%s%d tests failed%s\n", colors.Red(""), failureCount, colors.Reset(""))
	} else {
		fmt.Printf("%sAll tests passed! Great job!%s\n", colors.Green(""), colors.Reset(""))
	}
}

func main() {
	fmt.Println("Running unit tests for IsAnagram function:")
	runTests()
}
