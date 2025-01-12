package longest_consecutive

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

const (
	MIN_ARRAY_SIZE = 0
	MAX_ARRAY_SIZE = 1000
	MIN_VALUE      = -10
	MAX_VALUE      = 10
	TEST_CASES_NUM = 100
)

// Generate a random array within the problem constraints
func generateRandomArray() []int {
	size := rand.Intn(MAX_ARRAY_SIZE-MIN_ARRAY_SIZE+1) + MIN_ARRAY_SIZE
	result := make([]int, size)
	for i := range result {
		result[i] = rand.Intn(MAX_VALUE-MIN_VALUE+1) + MIN_VALUE
	}
	return result
}

// Generate multiple test cases
func generateFuzzyTestCases(numCases int) [][]int {
	var testCases [][]int
	for i := 0; i < numCases; i++ {
		testCases = append(testCases, generateRandomArray())
	}
	return testCases
}

// Benchmark the solution using random test cases
func BenchmarkLongestConsecutiveSequence(b *testing.B) {
	testCases := generateFuzzyTestCases(TEST_CASES_NUM)
	expectedResults := make([]int, TEST_CASES_NUM)

	// Precompute results for comparison
	for i, testCase := range testCases {
		expectedResults[i] = longestConsecutiveNeetcode(testCase)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i, tc := range testCases {
			actual := longestConsecutive(tc)
			if actual != expectedResults[i] {
				b.Errorf("Failed for test case %v", tc)
				b.Logf("Expected: %d, Got: %d", expectedResults[i], actual)
			}
		}
	}
}

func Convert(input string) []int {

	// Remove the surrounding brackets and trim spaces
	input = strings.Trim(input, "[]")
	// Split the string into parts based on spaces
	parts := strings.Fields(input)

	// Convert parts to integers
	result := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return make([]int, 0) // Return error if conversion fails
		}
		result[i] = num
	}
	return result
}

func BenchmarkLongestConsecutiveSequenceSingle(b *testing.B) {
	testCases := Convert("[1 5 10 7 -5]")

	result := longestConsecutiveNeetcode(testCases)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		actual := longestConsecutive(testCases)
		if actual != result {
			b.Errorf("Failed for test case %v", testCases)
			b.Logf("Expected: %d, Got: %d", result, actual)
		}
	}
}
