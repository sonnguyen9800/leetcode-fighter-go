package topkfreqelement

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

const (
	TEST_CASES_NUMBER = 1000
	MIN_ARRAY_SIZE    = 1
	MAX_ARRAY_SIZE    = 10000

	MIN_NUMBER = -1000
	MAX_NUMBER = 1000
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func generateTestCase() ([]int, int, []int) {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(MAX_ARRAY_SIZE-MIN_ARRAY_SIZE+1) + MIN_ARRAY_SIZE
	k := rand.Intn(min(n, 3)) + 1

	// Generate k distinct elements for the expected array
	var expected []int
	freq := make(map[int]int)
	uniqueFreq := make(map[int]bool)

	// Step 1: Generate k elements with unique high frequencies
	for len(expected) < k {
		val := rand.Intn(MAX_NUMBER-MIN_NUMBER+1) + MIN_NUMBER
		if _, exists := freq[val]; !exists {
			newFreq := len(expected) + n // Ensure large gaps to prevent overlaps
			for uniqueFreq[newFreq] {
				newFreq++
			}
			expected = append(expected, val)
			freq[val] = newFreq
			uniqueFreq[newFreq] = true
		}
	}

	// Step 2: Fill nums array with the generated frequencies
	var nums []int
	for val, count := range freq {
		for i := 0; i < count; i++ {
			nums = append(nums, val)
		}
	}

	// Step 3: Add random elements with strictly lower frequencies
	remainingSpace := n - len(nums)
	if remainingSpace > 0 {
		for len(nums) < n {
			val := rand.Intn(MAX_NUMBER-MIN_NUMBER+1) + MIN_NUMBER
			if _, exists := freq[val]; !exists {
				lowFreq := rand.Intn(min(4, len(expected))) + 1
				for i := 0; i < lowFreq && len(nums) < n; i++ {
					nums = append(nums, val)
				}
			}
		}
	}

	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

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

func BenchmarkTopKFreq(b *testing.B) {

	testCases := generateFuzzyTestCases(TEST_CASES_NUMBER)
	b.ResetTimer()
	// b.Run("Custom Solution", func(b *testing.B) {
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
	// b.ResetTimer()

	b.Run("Custom Solution - Optimal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tc := range testCases {
				nums := tc[0].([]int)
				k := tc[1].(int)
				expected := tc[2].([]int)
				got := TopKFreqElement_Optimal(nums, k)

				if !compareUnorderedSlices(got, expected) {
					b.Errorf("Failed: \n")
					b.Logf("nums: %v", nums)
					b.Logf("k: %v", k)

					b.Logf("Expected: %v", expected)
					b.Logf("Got: %v", got)

				}
			}
		}
	})
	b.ResetTimer()
	b.Run("Sorting - Neetocode Solution", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tc := range testCases {
				nums := tc[0].([]int)
				k := tc[1].(int)
				expected := tc[2].([]int)

				got := topKFrequent_sorting_neetcodesolution(nums, k)

				if !compareUnorderedSlices(got, expected) {
					b.Errorf("Failed: \n")
					b.Logf("nums: %v", nums)
					b.Logf("k: %v", k)

					b.Logf("Expected: %v", expected)
					b.Logf("Got: %v", got)

				}
			}
		}
	})
	b.ResetTimer()
	b.Run("Heap_(NeetcodeSolution)", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tc := range testCases {
				nums := tc[0].([]int)
				k := tc[1].(int)
				expected := tc[2].([]int)

				got := topKFrequent_heap_neetcode_solution(nums, k)

				if !compareUnorderedSlices(got, expected) {
					b.Errorf("Failed: \n")
					b.Logf("nums: %v", nums)
					b.Logf("k: %v", k)

					b.Logf("Expected: %v", expected)
					b.Logf("Got: %v", got)

				}
			}
		}
	})
	b.ResetTimer()
	b.Run("BucketSort(NeetcodeSolution)", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tc := range testCases {
				nums := tc[0].([]int)
				k := tc[1].(int)
				expected := tc[2].([]int)

				got := topKFrequent_bucket_sort(nums, k)

				if !compareUnorderedSlices(got, expected) {
					b.Errorf("Failed: \n")
					b.Logf("nums: %v", nums)
					b.Logf("k: %v", k)

					b.Logf("Expected: %v", expected)
					b.Logf("Got: %v", got)

				}
			}
		}
	})
	b.ResetTimer()
	b.Run("Using 3rd Party(Neetcode Solution)", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tc := range testCases {
				nums := tc[0].([]int)
				k := tc[1].(int)
				expected := tc[2].([]int)

				got := topKFrequent_neetcodesolution_3rdlibrary(nums, k)

				if !compareUnorderedSlices(got, expected) {
					b.Errorf("Failed: \n")
					b.Logf("nums: %v", nums)
					b.Logf("k: %v", k)

					b.Logf("Expected: %v", expected)
					b.Logf("Got: %v", got)

				}
			}
		}
	})
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
