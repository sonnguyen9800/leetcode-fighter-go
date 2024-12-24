package topkfreqelement

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
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

func topKFrequent_neetcodesolution(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	heap := priorityqueue.NewWith(func(a, b interface{}) int {
		freqA := a.([2]int)[0]
		freqB := b.([2]int)[0]
		return utils.IntComparator(freqA, freqB)
	})

	for num, freq := range count {
		heap.Enqueue([2]int{freq, num})
		if heap.Size() > k {
			heap.Dequeue()
		}
	}

	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		value, _ := heap.Dequeue()
		res[i] = value.([2]int)[1]
	}
	return res
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

func BenchmarkMySolution(b *testing.B) {
	b.ResetTimer()

	simple_nums := []int{1, 1, 1, 2, 2, 3, 88, 2, 2, 91, 2, 88}
	simple_k := 3
	simple_expected := []int{1, 2, 88}

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

	testCases := generateFuzzyTestCases(TEST_CASES_NUMBER)

	b.Run("Main", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tc := range testCases {
				nums := tc[0].([]int)
				k := tc[1].(int)
				expected := topKFrequent_neetcodesolution(nums, k)
				got := TopKFreqElement(nums, k)

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
