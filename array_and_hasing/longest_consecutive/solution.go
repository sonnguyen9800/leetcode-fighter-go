package longest_consecutive

func longestConsecutive(nums []int) int {

	hash := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = true
	}
	greatestConsCount := 0
	for i := 0; i < len(nums); i++ {
		if hash[nums[i]-1] == false {
			lenght := 1
			for hash[nums[i]+lenght] == true {
				lenght++

			}
			if lenght > greatestConsCount {
				greatestConsCount = lenght
			}
		}
	}
	return greatestConsCount
}

func longestConsecutiveOnePassB(nums []int) int {

	hash := make(map[int]int)

	greatestConsCount := 0
	for _, num := range nums {
		if hash[num] == 0 {
			next := hash[num+1]
			previous := hash[num-1]

			//if previous == 0 && next == 0 {
			//	hash[num] = 1
			//}

			if next != 0 && previous == 0 {
				prev_series_length := hash[num+1] // this also MUST containt length of the consecutive
				last_value_in_series := num + prev_series_length
				hash[last_value_in_series] = hash[last_value_in_series] + 1
				hash[num] = hash[last_value_in_series] // Also update "new" last index to the length of series
			}

			if previous != 0 && next == 0 {
				// if num is new leading number of the series, with no "follower"
				last_leading_value := hash[num-1] // this also MUST containt length of the consecutive
				hash[num] = last_leading_value + 1
				hash[num-last_leading_value] = last_leading_value + 1 // this supposed to be the "first" value in the series ("lower-bound")

			}

			if (previous != 0 && next != 0) || (previous == 0 && next == 0) {
				last_leading_value := hash[num-1]  // the leading of a 1st series (upper bound)
				first_leading_value := hash[num+1] // the lower bound of a series

				new_series_length := last_leading_value + first_leading_value + 1
				hash[num] = new_series_length
				hash[num-last_leading_value] = new_series_length
				hash[num+first_leading_value] = new_series_length
			}

		}
	}
	for _, num := range hash {
		if num > greatestConsCount {
			greatestConsCount = num
		}
	}
	return greatestConsCount
}
func longestConsecutiveOnePass(nums []int) int {

	hash := make(map[int]int)

	greatestConsCount := 0
	consecutiveCount := 0
	for _, num := range nums {
		if hash[num] == 0 {
			next := hash[num+1]
			previous := hash[num-1]

			if previous == 0 && next == 0 {
				hash[num] = 1
				consecutiveCount = hash[num]
			}

			if next != 0 && previous == 0 {
				hash[num+1] = hash[num+1] + 1
				consecutiveCount = hash[num+1]
				hash[num] = hash[num+1]
			}

			if previous != 0 && next == 0 {
				hash[num] = hash[num-1] + 1
				consecutiveCount = hash[num]
				hash[num-1] = hash[num]
				hash[num-hash[num]+1] = hash[num]
			}

			if previous != 0 && next != 0 {
				hash[num+1] = hash[num+1] + hash[num-1] + 1
				hash[num] = hash[num+1]
				hash[num-1] = hash[num+1]
				if hash[num-hash[num]+1] != 0 {
					hash[num-hash[num]+1] = hash[num]

				}

				consecutiveCount = hash[num+1]
			}

			if consecutiveCount > greatestConsCount {
				greatestConsCount = consecutiveCount
			}
		}
	}
	return greatestConsCount
}
