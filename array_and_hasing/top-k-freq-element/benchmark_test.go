package topkfreqelement

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

const (
	TEST_CASES_NUMBER = 1
	MIN_ARRAY_SIZE    = 1
	MAX_ARRAY_SIZE    = 10000
	MIN_NUMBER        = -1000
	MAX_NUMBER        = 1000
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Generates a random test case for the problem, including expected output
func generateTestCase() ([]int, int, []int) {

	n := rand.Intn(MAX_ARRAY_SIZE-MIN_ARRAY_SIZE+1) + MIN_ARRAY_SIZE
	var nums []int
	freq := make(map[int]int)

	for i := 0; i < n; i++ {
		val := rand.Intn(MAX_NUMBER-MIN_NUMBER+1) + MIN_NUMBER
		nums = append(nums, val)
		freq[val]++
	}
	distinctCount := len(freq)
	k := 1
	if distinctCount > 0 {
		k = rand.Intn(distinctCount) + 1
	}

	// Calculate the k most frequent elements
	type kv struct {
		key   int
		value int
	}
	var freqSlice []kv
	for key, value := range freq {
		freqSlice = append(freqSlice, kv{key, value})
	}

	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].value > freqSlice[j].value
	})

	var expected []int
	for i := 0; i < min(k, len(freqSlice)); i++ {
		expected = append(expected, freqSlice[i].key)
	}

	return nums, k, expected
}

func generateFuzzyTestCases(numCases int) [][3]interface{} {
	var testCases [][3]interface{}
	for i := 0; i < numCases; i++ {
		nums, k, expected := generateTestCase()
		testCases = append(testCases, [3]interface{}{nums, k, expected})
	}
	return testCases
}

func BenchmarkMySolution(b *testing.B) {
	b.ResetTimer()

	simple_nums := []int{1, 1, 1, 2, 2, 3}
	simple_k := 2
	simple_expected := []int{1, 2}

	b.Run("MySolution Simple", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			got := TopKFreqElement(simple_nums, simple_k)
			if !compareUnorderedSlices(got, simple_expected) {
				b.Errorf("Failed: \n")
				b.Logf("nums: %v", simple_nums)
				b.Logf("k: %v", simple_k)

				b.Logf("Expected: %v", simple_expected)
				b.Logf("Got: %v", got)

			}
		}
	})
	// 	testCases := generateFuzzyTestCases(TEST_CASES_NUMBER)

	// b.Run("MySolution", func(b *testing.B) {
	// 	for n := 0; n < b.N; n++ {
	// 		for _, tc := range testCases {
	// 			nums := tc[0].([]int)
	// 			k := tc[1].(int)
	// 			expected := tc[2].([]int)
	// 			got := TopKFreqElement(nums, k)
	// 			if !compareUnorderedSlices(got, expected) {
	// 				b.Errorf("Failed: \n")
	// 				b.Logf("nums: %v", nums)
	// 				b.Logf("k: %v", k)

	// 				b.Logf("Expected: %v", expected)
	// 				b.Logf("Got: %v", got)

	// 			}
	// 		}
	// 	}
	// })

	b.ResetTimer()
}

func compareUnorderedSlices(a, b []int) bool {
	mapA := make(map[int]int)
	mapB := make(map[int]int)

	for _, v := range a {
		mapA[v]++
	}
	for _, v := range b {
		mapB[v]++
	}

	return reflect.DeepEqual(mapA, mapB)
}
