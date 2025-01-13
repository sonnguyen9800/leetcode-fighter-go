package longest_consecutive

import "testing"

func longestConsecutiveNeetcode(nums []int) int {
	mp := make(map[int]int)
	res := 0

	for i, num := range nums {
		if mp[num] == 0 {
			left := mp[num-1]
			right := mp[num+1]
			sum := left + right + 1
			mp[num] = sum
			mp[num-left] = sum
			mp[num+right] = sum
			if sum > res {
				print("Position %d", i)
				res = sum
			}
		}
	}
	return res
}

func longestConsecutiveNeetcodeB(nums []int, b *testing.B) int {
	mp := make(map[int]int)
	res := 0

	for i, num := range nums {
		if mp[num] == 0 {
			left := mp[num-1]
			right := mp[num+1]
			sum := left + right + 1
			mp[num] = sum
			mp[num-left] = sum
			mp[num+right] = sum
			if sum > res {
				print("Position %d", i)
				res = sum
			}
		}
	}
	//b.Logf("Data [0] %d", mp)
	return res
}
