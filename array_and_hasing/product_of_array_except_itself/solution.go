package product_of_array_except_itself

func productExceptSelf_Solution1(nums []int) []int {
	totalProduct := 1
	lenght := len(nums)
	result := make([]int, lenght)

	for i := 0; i < lenght; i++ {
		totalProduct = totalProduct * nums[i]
	}
	for i := 0; i < lenght; i++ {
		result[i] = totalProduct / nums[i]
	}
	return result
}

func productExceptSelf_Solution2(nums []int) []int {
	totalProduct := 1
	lenght := len(nums)
	result1 := make([]int, lenght)
	result2 := make([]int, lenght)
	result3 := make([]int, lenght)

	for i := 0; i < lenght; i++ {
		if i == 0 {
			result1[i] = 1
			continue
		}
		result1[i] = totalProduct * nums[i-1]
		totalProduct = totalProduct * nums[i-1]
	}
	totalProduct = 1
	for i := lenght - 1; i >= 0; i-- {
		if i == lenght-1 {
			result2[i] = 1
			continue
		}
		result2[i] = totalProduct * nums[i+1]
		totalProduct = totalProduct * nums[i+1]

	}

	for i := 0; i < lenght; i++ {
		result3[i] = result1[i] * result2[i]
	}
	return result3
}
