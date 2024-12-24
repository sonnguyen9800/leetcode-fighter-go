package topkfreqelement

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

/// Optimal to: on the loop to craete freqMapByCount, we can also find the highest frequency there
///

func TopKFreqElement_Optimal(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num] += 1
	}

	// Creating a map of frequency of each frequency
	freqMapByCount := make(map[int][]int)
	highest_frequency := 0

	for key, freqValue := range freqMap {
		freqMapByCount[freqValue] = append(freqMapByCount[freqValue], key)
		if freqValue > highest_frequency {
			highest_frequency = freqValue
		}
	}

	output_array := []int{}

	for i := highest_frequency; i >= 0; i-- {
		for _, num := range freqMapByCount[i] {
			output_array = append(output_array, num)
			if len(output_array) == k {
				return output_array
			}
		}

	}
	return output_array
}
