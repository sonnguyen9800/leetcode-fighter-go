---
title: "Leetcode Fighter: Group Anagrams"
author: "Son Nguyen Hoang"
abstract: "Documents and study note for classic leetcode problem: Anagram Groups"
date: 2024-11-26T16:00:53+07:00
draft: false
language: "English"
tags: ["golang", "leetcode", "leetcode-fighter"]
thumbnail: thumbnail.jpg
comment: true
ShowSummary: true
---

![alt text](image.png)

## Group Anagrams

### Problem Statement

Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

Example 1:

>Input: strs = ["act","pots","tops","cat","stop","hat"]
>Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

Example 2:

>Input: strs = ["x"]
>Output: [["x"]]

Example 3:

>Input: strs = [""]
>Output: [[""]]

#### Constraints:
```
Constraints:

1 <= strs.length <= 1000.
0 <= strs[i].length <= 100
strs[i] is made up of lowercase English letters.
```

### A short analysis

The problem is not hard. The first instinct of mine would be to reuse solution in "IsAnagram" and adapt to this problem.
The change would be to make a hashmap with _string_ to be a key. We would love to every string in the list. Check if the
hashmap has the key which is an anagram of that string. If exist then append that string into the map. If not then create 
a new string.

The funnier part when solving this solution would be to write a fuzzy test cases generator to mass-produce test cases for 
benchmarking. 

### Solution:

#### First Attempt

```
func CheckIfAnagramExist(string2 string, map_input map[string][]string) {
	exist_key := false
	for key, _ := range map_input {
		if (is_anagram.IsAnagram(is_anagram.InputString{
			S1: string2,
			S2: key,
		})) {
			map_input[key] = append(map_input[key], string2)
			exist_key = true
		}
	}
	if !exist_key {
		map_input[string2] = []string{string2}
	}

}

func GroupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		CheckIfAnagramExist(str, anagramMap)
	}

	data := [][]string{}

	for _, strs_value := range anagramMap {
		array_str := []string{}
		for _, str := range strs_value {
			array_str = append(array_str, str)
		}
		data = append(data, array_str)
	}

	return data
}
```
Although I am confident that this time complexity is O(m * n), I also have to admit that neetcode's solution is much more clean
and faster. Benchmarking the two with N = 1000 and N = 10,000 proved that my solution is only 60% as performative as Neetcode's

```
goos: windows
goarch: amd64
pkg: leetcode-fighter-go/array_and_hasing/anagram_groups
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkGroupAnagrams
BenchmarkGroupAnagrams/MySolution_
BenchmarkGroupAnagrams/MySolution_-16              20464             56304 ns/op

BenchmarkGroupAnagrams/Neetcode_
BenchmarkGroupAnagrams/Neetcode_-16                36330             33041 ns/op

BenchmarkGroupAnagrams/Test_10000-16               2085             563571 ns/op

BenchmarkGroupAnagrams/Neetcode_
BenchmarkGroupAnagrams/Neetcode_10000_-16          3574             339131 ns/op
```

### Neetcode Solution


```
func groupAnagramsNeetCodeSolution(strs []string) [][]string {
	res := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		res[count] = append(res[count], s)
	}

	var result [][]string
	for _, group := range res {
		result = append(result, group)
	}
	return result
}
```
Again, he abused the assumption that all characters in string are English alphabet. His overall strategy is the same as mine
(proved when my algorithm scales as good as him), but he did some micro-optimization that briliantly require less code, but 50% 
faster.
- Instead of a `string` as key (like mine), he used `array of rune` as key
- Then instead check if any aragram of input string is the same as the key (as mine), he loops through string (to be `rune`) and create
array from it.
- This array is reused as `key` for the `hashmap`

The key interesting point is that array as 'key' can be essily access by value. It turns out that every comparable value can be
used as `key`. Also, every array of comparabled value can be comparabled: [Source](https://go.dev/ref/spec#Comparison_operators)

```
Array types are comparable if their array element types are comparable. Two array values are equal 
if their corresponding element values are equal. The elements are compared in ascending index order, 
and comparison stops as soon as two element values differ (or all elements have been compared).
```
Very brilliant! This inspired me to make a solution of mine (work with generic character instead)

### New Solution

```
func GroupAnagramNew(strs []string) [][]string {
	res := make(map[string][]string)
	sortified := func(str string) string {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		return string(runes)
	}
	for _, str := range strs {
		sorted_string := sortified(str)
		res[sorted_string] = append(res[sorted_string], str)
	}
	data := [][]string{}
	for _, strs_value := range res {
		data = append(data, strs_value)
	}
	return data
}
```

What changes:
- Sort each the `string` to check if it's similar to the 'key'
- If yes, append to the hashmap, if not, create a new key

Pretty significant change as long as I concern. However, The truth is that its performance is not really ... a boost from my original version

```
goos: windows
goarch: amd64
pkg: leetcode-fighter-go/array_and_hasing/anagram_groups
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkGroupAnagrams
BenchmarkGroupAnagrams/First_Solution:_N=1000
BenchmarkGroupAnagrams/First_Solution:_N=1000-16                   20601    59440 ns/op
BenchmarkGroupAnagrams/Neetcode:_N=1000
BenchmarkGroupAnagrams/Neetcode:_N=1000-16                         35193    33629 ns/op
BenchmarkGroupAnagrams/New_Solution:_N=1000
BenchmarkGroupAnagrams/New_Solution:_N=1000-16                     19917    60317 ns/op
PASS
```
With N=10,000

```
BenchmarkGroupAnagrams
BenchmarkGroupAnagrams/First_Solution:_N=10000
BenchmarkGroupAnagrams/First_Solution:_N=10000-16                   2030    568944 ns/op
BenchmarkGroupAnagrams/Neetcode:_N=10000
BenchmarkGroupAnagrams/Neetcode:_N=10000-16                         3373    342457 ns/op
BenchmarkGroupAnagrams/New_Solution:_N=10000
BenchmarkGroupAnagrams/New_Solution:_N=10000-16                     2000    590706 ns/op

```

It's hard to believe that with such optimization. Nothing really change at all. It's seem like 
the core of difference is the smart use of array (len: 26) is perfectly fit given that the input includes only English-alphabet character only

# Conclusion
- It's funny and remarkable to realize that a small change in given scenario can lead to dramatic difference!