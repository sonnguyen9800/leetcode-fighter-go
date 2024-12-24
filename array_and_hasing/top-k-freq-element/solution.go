package topkfreqelement

import (
	"sort"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

func TopKFreqElement(nums []int, k int) []int {
	length := len(nums)
	if k > length {
		return []int{}
	}

	// Creating a map of frequency of each element
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num] += 1
	}

	// Creating a map of frequency of each frequency
	freqMapByCount := make(map[int][]int)
	for key, freqValue := range freqMap {
		if freqMapByCount[freqValue] == nil {
			freqMapByCount[freqValue] = []int{}
		}
		freqMapByCount[freqValue] = append(freqMapByCount[freqValue], key)
	}

	output_array := []int{}

	for i := 9999; i >= 0; i-- {

		if len(freqMapByCount[i]) == 0 {
			continue
		}
		current_length := len(output_array)
		if current_length >= k {
			break
		}
		next_length := len(freqMapByCount[i])

		if (next_length + current_length) > k {
			var amount int
			if k > current_length {
				amount = k - current_length
			} else {
				amount = current_length - k
			}
			output_array = append(output_array, freqMapByCount[i][0:amount]...)
			break
		} else {
			output_array = append(output_array, freqMapByCount[i]...)
		}
	}
	return output_array

}

func topKFrequent_sorting_neetcodesolution(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	arr := make([][2]int, 0, len(count))
	for num, cnt := range count {
		arr = append(arr, [2]int{cnt, num})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i][1]
	}
	return res
}

func topKFrequent_heap_neetcode_solution(nums []int, k int) []int {
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

func topKFrequent_bucket_sort(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		count[num]++
	}
	for num, cnt := range count {
		freq[cnt] = append(freq[cnt], num)
	}

	res := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
