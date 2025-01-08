package product_of_array_except_itself

import (
	"math/rand"
	"testing"
)

const (
	MIN_ARRAY_SIZE = 2
	MAX_ARRAY_SIZE = 5
	MIN_VALUE      = -20
	MAX_VALUE      = 20
	TEST_CASES_NUM = 1
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

// CompareArrays checks if two slices are equal
func CompareArrays(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// Benchmark the solution using random test cases
func BenchmarkProductExceptSelf(b *testing.B) {
	testCases := generateFuzzyTestCases(TEST_CASES_NUM)
	result := make([][]int, TEST_CASES_NUM)
	for i, testCase := range testCases {
		result[i] = productExceptSelf_Neetcode(testCase)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i, tc := range testCases {
			actual := productExceptSelf_Solution2(tc)
			if !CompareArrays(actual, result[i]) {
				b.Errorf("Failed:\n")
				b.Logf("Origin: %d \n", tc)
				b.Logf("Expected: %d \n", result[i])
				b.Logf("Actual: %d \n", actual)

			}

		}
	}
}
