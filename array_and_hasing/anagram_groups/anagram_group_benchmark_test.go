package anagram_groups

import (
	"fmt"
	"sort"
	"testing"
)

var N = 1000
var MinNumGroup = 1
var MaxNumGroup = 2

var MinGroupSize = 1
var MaxGroupSize = 10
var RanomNoiseGroupNumber = 10

func normalizeArray(arr [][]string) [][]string {
	// Sort each sublist
	for i := range arr {
		sort.Strings(arr[i])
	}
	// Sort the outer list
	sort.Slice(arr, func(i, j int) bool {
		return compareSlices(arr[i], arr[j]) < 0
	})
	return arr
}

// compareSlices compares two string slices lexicographically.
func compareSlices(a, b []string) int {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}
	return len(a) - len(b)
}

// areArraysEqual checks if two arrays of string arrays are equal.
func areArraysEqual(arr1, arr2 [][]string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	arr1 = normalizeArray(arr1)
	arr2 = normalizeArray(arr2)
	for i := range arr1 {
		if !equalSlices(arr1[i], arr2[i]) {
			return false
		}
	}
	return true
}

// equalSlices checks if two string slices are equal.
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkGroupAnagrams(b *testing.B) {

	input_lst := make([][]string, N)
	expected_lst := make([][][]string, N)
	for i := 0; i < N; i++ {
		input, expected := fuzzyTestCaseGeneratorAdv(
			MinNumGroup, MaxNumGroup,
			MinGroupSize, MaxGroupSize,
			RanomNoiseGroupNumber)

		input_lst = append(input_lst, input)
		expected_lst = append(expected_lst, expected)

	}

	b.Run(fmt.Sprint("Test "), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				actual := GroupAnagrams(input_lst[i])
				if !areArraysEqual(actual, expected_lst[i]) {
					b.Errorf("Input: %v", input_lst[i])
					b.Errorf("Actual %v \n Expected %v", actual, expected_lst[i])
				}
			}
		}

	})

}
