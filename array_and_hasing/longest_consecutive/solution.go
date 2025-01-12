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
