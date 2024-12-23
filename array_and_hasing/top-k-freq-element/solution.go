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
			amount := k - current_length
			output_array = append(output_array, freqMapByCount[i][amount:]...)
			break
		} else {
			output_array = append(output_array, freqMapByCount[i]...)
		}
	}
	return output_array

}