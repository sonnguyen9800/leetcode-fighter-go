package main

import (
	"fmt"
	"leetcode-fighter-go/goland-slib/colors" // Replace with your actual module name
)

// Unit test function
func runTests() {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		// No duplicates
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{10, 20, 30, 40, 50, 60}, false},
		{[]int{100, 200, 300, 400}, false},
		{[]int{5, 6, 7, 8, 9, 10}, false},
		{[]int{2, 4, 6, 8, 10}, false},

		// Duplicates
		{[]int{1, 2, 3, 4, 5, 3}, true},
		{[]int{1, 1}, true},
		{[]int{1, 2, 2, 3, 4}, true},
		{[]int{10, 20, 30, 30, 40}, true},
		{[]int{5, 5, 5, 5}, true},
		{[]int{-1, -2, -3, -1}, true},
		{[]int{0, 0}, true},
		{[]int{1, 2, 3, 1}, true},
		{[]int{3, 3, 3}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 7}, true},

		// Edge cases
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{1000000, -1000000}, false},
		{[]int{1000000, 1000000}, true},
		{[]int{0}, false},

		// Large input with duplicates
		{append(make([]int, 10000), 1), true},
		{append(make([]int, 10000), 0, 0), true},
		{append(make([]int, 10000), 9999), true},

		// Negative numbers
		{[]int{-1, -2, -3, -2}, true},
		{[]int{-1, -2, -3}, false},
		{[]int{-10, -20, -30, -40, -20}, true},
		{[]int{-5, -6, -7, -8}, false},
		{[]int{-1, -1}, true},

		// Mixed numbers
		{[]int{1, -1, 2, -2, 3, -3}, false},
		{[]int{1, -1, 2, -2, 2}, true},
		{[]int{-10, 10, -20, 20, -10}, true},
		{[]int{5, 6, 5, 7}, true},
		{[]int{-100, 100, -100, 200}, true},

		// Repeated patterns
		{[]int{1, 2, 3, 1, 2, 3}, true},
		{[]int{4, 5, 6, 4, 5, 6}, true},
		{[]int{10, 20, 10, 30, 40}, true},
		{[]int{0, 1, 2, 0, 1, 2}, true},
		{[]int{5, 5, 10, 10, 15}, true},

		// Random cases
		{[]int{9, 8, 7, 6, 5, 4, 4}, true},
		{[]int{100, 101, 102, 103, 104, 100}, true},
		{[]int{50, 51, 52, 53, 54, 55}, false},
		{[]int{2, 3, 4, 5, 6, 7, 2}, true},
		{[]int{9, 8, 7, 6, 5}, false},

		// Larger datasets
		{make([]int, 20000), true},                 // Large array with no duplicates
		{append(make([]int, 19999), 42), true},     // Large array with duplicate
		{append(make([]int, 20000), 100), true},    // Large array with duplicate
		{append(make([]int, 19999), -1, -1), true}, // Large array with negative duplicates
		{append(make([]int, 20000), 0), true},      // Large array with duplicate zero
	}

	successCount := 0
	failureCount := 0

	for i, testCase := range testCases {
		result := containsDuplicate(testCase.nums)
		if result != testCase.expected {
			failureCount++
			fmt.Printf("%sTest Case %d Failed: Expected %v, Got %v\n", colors.Red(""), i+1, testCase.expected, result)
			fmt.Printf("Input (Length %s): %v%s \n", len(testCase.nums),
				testCase.nums, colors.Reset(""))

			fmt.Printf("Before Input: %v%s\n", testCases[i-1].nums, colors.Reset(""))

		} else {
			successCount++
			fmt.Printf("%sTest Case %d Passed!%s\n", colors.Green(""), i+1, colors.Reset(""))
		}
	}

	// Summary with fancy output
	totalTests := len(testCases)
	fmt.Printf("\n%s%sSummary:%s %s%d/%d tests passed%s (%.2f%% success rate)\n", colors.Bold(""), colors.Green(""), colors.Reset(""), colors.Bold(""), successCount, totalTests, colors.Reset(""), float64(successCount)/float64(totalTests)*100)
	if failureCount > 0 {
		fmt.Printf("%s%d tests failed%s\n", colors.Red(""), failureCount, colors.Reset(""))
	} else {
		fmt.Printf("%sAll tests passed! Great job!%s\n", colors.Green(""), colors.Reset(""))
	}
}
func main() {
	fmt.Println("Running unit tests for containsDuplicate function:")
	runTests()
}
