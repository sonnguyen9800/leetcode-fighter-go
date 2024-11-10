# Is Anagram

## A short analysis

This problem is fairly simple, I would quote the original problem from neetcode.io below. Link: https://neetcode.io/problems/is-anagram

Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

> Given two strings `s` and `t`, return `<span class="token boolean">true</span>` if the two strings are anagrams of each other, otherwise return `<span class="token boolean">false</span>`.
>
> An **anagram** is a string that contains the exact same characters as another string, but the order of the characters can be different.

### Example:

```
Input: s = "racecar", t = "carrace" Output: true
Example 2:Input: s = "jar", t = "jam" Output: false
```

### Constraints:

s and t consist of lowercase English letters.

### Solution:

#### First Attempt

1. Compare length of two strings, return false immediately if the lengths of the two are not equal
2. Convert two strings to an array of rune
3. Creating dictionary (map in go) of each rune, named `map1` and `map2`. Each map has `key` to be the **rune** and each `key's value` is the **frequency** of each rune in each string
4. Compare key-value between *map1* and *map2*

### Second Attempt

It turns out that we literally don't have to make an array of rune. Instead in Go we can do interate directly through the `string`

```
	for i := 0; i < len(first_str); i++ { // first_str is a string
		print(first_str[i]) // print rune directly
	}
```

So we can optimize the code to be like

```
func IsAnagram(first_str, second_str string) bool {
	if len(first_str) != len(second_str) {
		return false
	}

	runeMap_first, runMap_second := make(map[rune]int), make(map[rune]int)

	for i := 0; i < len(first_str); i++ {
		runeMap_first[rune(first_str[i])]++
		runMap_second[rune(second_str[i])]++
	}

	for currentRune, value := range runeMap_first {
		value_2, exist := runMap_second[currentRune]
		if !exist || value != value_2 {
			return false
		}
	}
	return true
}
```

Which is quite convenient. But this is not optimal *yet*

### Third Attempt (modified from Neetcode's solution)

His (neetcode) solution is given below

```

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := [26]int{}
    for i := 0; i < len(s); i++ {
        count[s[i]-'a']++
        count[t[i]-'a']--
    }

    for _, val := range count {
        if val != 0 {
            return false
        }
    }
    return true
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


