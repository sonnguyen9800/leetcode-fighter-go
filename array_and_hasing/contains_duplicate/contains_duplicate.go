package main

func containsDuplicate(nums []int) bool {

	dict := make(map[int]int)

	for _, num := range nums {
		dict[num] = dict[num] + 1
		if dict[num] > 1 {
			return true
		}
	}
	return false
}
