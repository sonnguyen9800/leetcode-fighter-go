# Two Sum

## Problem Statement

Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target and i != j.

You may assume that every input has exactly one pair of indices i and j that satisfy the condition.

Return the answer with the smaller index first.

Example 1:

>Input:
>nums = [3,4,5,6], target = 7
>Output: [0,1]
>Explanation: nums[0] + nums[1] == 7, so we return [0, 1].

Example 2:

>Input: nums = [4,5,6], target = 10
>Output: [0,2]
>Example 3:

Input: nums = [5,5], target = 10

Output: [0,1]
### Constraints:

2 <= nums.length <= 1000
-10,000,000 <= nums[i] <= 10,000,000
-10,000,000 <= target <= 10,000,000

## A short analysis

The problem is not hard. The first instinct of mine would be create a 2d array then craft the matrix of sum of each then 
simply look for the wanted indices. When I started coding, I reliaze that a quick nested for loop is enough to Bruteforce 
and find the wanted indicies.

After finish the solution, I looked for Neetcode advice (his video) and immediately what he meant on the "optimized" solution. 
Which doesn't require the nested loop but one loop only. This is done by using a hashmap and look for "wanted" index. 

However, when evaluating the result, (benchmark), something weird happens.

### Solution:

#### First Attempt

```
func TwoSumBruteForce(input TwoIntergersInput) []int {
	num_array := input.ArrayInt
	target := input.Target

	for i := 0; i < len(num_array); i++ {
		for j := i + 0; j < len(num_array); j++ {
			if i == j {
				continue
			}
			if num_array[i]+num_array[j] == target {
				return []int{i, j}
			}

		}
	}

	return []int{0, 0}
}
```
The bruteforce approach is simple. Simply brute force and make sum between each index in the array until you find the 
wanted pair.

### Second Attempt: OnePass Method

My first draft on OnePass based on quick advice of Neetcode on Youtube looked like this

```
func TwoSumOnePass(input TwoIntergersInput) []int {
	dict_input := make(map[int]int)
	for i := 0; i < len(input.ArrayInt); i++ {
		dict_input[input.ArrayInt[i]] = i
	}
	for i := 0; i < len(input.ArrayInt); i++ {
		value, exists := dict_input[input.Target-input.ArrayInt[i]]

		if value != i && exists {
			return []int{value, i}
		}
	}

	return []int{0, 0}
}
```
The code run through the testcases without problem. However, performance wise, it doesn't compete with Brute force. The 
benchmark analysis will be discussed below

### Third Attempt OnePass V2

It's obvious that dynamically change size of the `map` cause serious issue. So this time I initialize the `map` first.
```
func TwoSumOnePassVer2(input TwoIntergersInput) []int {
	dict_input := make(map[int]int, len(input.ArrayInt)) // Fixed size for the map
	for i := 0; i < len(input.ArrayInt); i++ {
		dict_input[input.ArrayInt[i]] = i
	}
    // ... The rest
}
```

### Third Attempt OnePass V3

The initialization of map can be done inside the loop instead. So I remove the 1st loop from the function.
```
func TwoSumOnePassVer3(input TwoIntergersInput) []int {
	dict_input := make(map[int]int, len(input.ArrayInt))

    // No more loop here
    
	for i := 0; i < len(input.ArrayInt); i++ {
		value, exists := dict_input[input.Target-input.ArrayInt[i]]

		if value != i && exists {
			return []int{value, i}
		}
		dict_input[input.ArrayInt[i]] = i
	}

	return nil
}
```

### Third Attempt OnePass V3

The initialization of map can be done inside the loop instead. So I remove the 1st loop from the function.
```
func TwoSumOnePassVer3(input TwoIntergersInput) []int {
	dict_input := make(map[int]int, len(input.ArrayInt))

    // No more loop here
    
	for i := 0; i < len(input.ArrayInt); i++ {
		value, exists := dict_input[input.Target-input.ArrayInt[i]]

		if value != i && exists {
			return []int{value, i}
		}
		dict_input[input.ArrayInt[i]] = i
	}

	return nil
}
```

Which is quite smart, as the beginning, I forgot that one of the constrats is that all letters are English. Given that prequisite, he

- Create an array of 26 element, (represent 26 Alphabet character).
- Each array contains a numeric integer.
- This integer is the `frequency` of each letter in both string, albeit there is a slight modification here. This `frequency` is not normal duplicate but act like the "difference" between frequency of a letter between first string and second string.
- E.g: If letter 'a' appear 9 times in first string and 4 times in second string. That mean the "difference" of frequency between string #1 and string #2 would be 9 - 4 = 5.
- Note that this difference can be negative as well.
- If this "difference" is 0, then string #1 has the same number of letter 'a' as string #2
- So, what the final loop does is to loop through the array of 26 elements, and check if any value of its is not 0. If such value existed, the two string (#1 and #2) cannot be anagram.


### Improvement

This is optimal solution if the input is Alphabet, lower-case character. But I don't want this two constant. What I improve is to

- Convert string to lowercase first
- Use a map instead of 26-character array (to support non-english letters)

```
unc IsAnagramOptimal(first_str, second_str string) bool {
	first_str := strings.ToLower(inputString.S1)
	second_str := strings.ToLower(inputString.S2)
	if len(first_str) != len(second_str) {
		return false
	}
	if len(first_str) != len(second_str) {
		return false
	}

	runeMap := make(map[rune]int)
	for i := 0; i < len(first_str); i++ {
		runeMap[rune(first_str[i])] += 1
		runeMap[rune(second_str[i])] -= 1
	}
	for _, value := range runeMap {
		if value != 0 {
			return false
		}
	}
	return true
}
```

In the code, I did a slight modification: to make the input become a custom struct type, which help us to write better testcases.


### Evaluation

The difference between optimal solution and first solution is that we use less than one `map`. Thus, we reduce usage of memory.

I did some test between two solution and provide the table below for you to see the difference


| Solution         | N Count = 101     | Test #1 | Test #2 | Test #3 | Test #4 | Test #5 | Average |
|------------------|-------------------|---------|---------|---------|---------|---------|---------|
| Normal Solution  | Time Cost    (ms) | 8.83    | 8.99    | 7.78    | 8.45    | 9.03    | 8.595   |
|                  | Memory (byte)     | 79160   | 79648   | 79712   | 80000   | 79856   | 79678   |
| Optimal Solution | Time Cost         | 8.38    | 8.17    | 7.84    | 9.29    | 8.56    | 8.448   |
|                  | Memory (byte)     | 73816   | 73888   | 73208   | 73760   | 73824   | 73699   |

This indicates that the optimal solution use **8%** less memory than the normal solution. This is because we use less than one map in the source code.


