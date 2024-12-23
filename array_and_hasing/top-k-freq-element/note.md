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
### Third Attempt OnePass V4

Finally, I realize that this time I don't even need to prefix the size of the map! Because the map is modified in the loop
itself!
```
func TwoSumOnePassVer4(input TwoIntergersInput) []int {
	dict_input := make(map[int]int)
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

That's my final solution, which is relatively similar to Neetcode (his is slightly different)

### Testing Scenario

- We create a fuzzy-testcase generator to create random test cases
- The algorithm will not run onced but run in a set of MANY set of inputs (array). The default is 100 set of input.
- This same set of testcases (100 testcases) is fetched into the function to evaluate the performance.

Beware, the result might shock you!

### Evaluation

#### A Paradox
The below benchmark is run with the size of array is 50,000 index. Note that I added Neetcode's optimal solution below and 
name it V5 (with slight modification to fit my framework)

>os: windows
goarch: amd64
pkg: leetcode-fighter-go/array_and_hasing/two_integers_sum
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkTwoSum/Bruteforce-16         	     669	   2019879 ns/op	    1600 B/op	     100 allocs/op
BenchmarkTwoSum/OnePassV1-16          	       9	 115684656 ns/op	68717328 B/op	   27720 allocs/op
BenchmarkTwoSum/OnePassV2-16          	      12	  93940642 ns/op	125341753 B/op	     309 allocs/op
BenchmarkTwoSum/OnePassV3-16          	     225	   6586528 ns/op	125341655 B/op	     300 allocs/op
BenchmarkTwoSum/OnePassV4-16          	    1218	    939956 ns/op	 1239443 B/op	    2146 allocs/op
BenchmarkTwoSum/OnePassV5-16          	    1266	    921025 ns/op	 1239625 B/op	    2147 allocs/op
PASS
ok  	leetcode-fighter-go/array_and_hasing/two_integers_sum	9.593s


What crazy from the benchmark is that:
- OnePass solution Ver 1 - 3 is **SLOWER** than `Bruteforce`
- It is even crazy that with N (lenght of input array) greater by magnitude, it's look like this trend never change! 

More benchmark is below

#### With N = 100,000
>goos: windows
goarch: amd64
pkg: leetcode-fighter-go/array_and_hasing/two_integers_sum
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkTwoSum/Bruteforce-16         	     242	   5108621 ns/op	    1600 B/op	     100 allocs/op
BenchmarkTwoSum/OnePassV1-16          	       5	 218696800 ns/op	68722696 B/op	   27717 allocs/op
BenchmarkTwoSum/OnePassV2-16          	       8	 156441975 ns/op	250679224 B/op	     300 allocs/op
BenchmarkTwoSum/OnePassV3-16          	      80	  13116000 ns/op	250679239 B/op	     300 allocs/op
BenchmarkTwoSum/OnePassV4-16          	     991	   1036371 ns/op	 1415327 B/op	    2308 allocs/op
BenchmarkTwoSum/OnePassV5-16          	    1141	   1042960 ns/op	 1415464 B/op	    2310 allocs/op
PASS
ok  	leetcode-fighter-go/array_and_hasing/two_integers_sum	11.028s

#### With N = 10,000,000

>goos: windows
goarch: amd64
pkg: leetcode-fighter-go/array_and_hasing/two_integers_sum
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkTwoSum/Bruteforce-16         	       2	 503096350 ns/op	    1604 B/op	     100 allocs/op
BenchmarkTwoSum/OnePassV1-16          	       1	18888793900 ns/op	68730208 B/op	   27807 allocs/op
BenchmarkTwoSum/OnePassV2-16          	       1	14242295500 ns/op	32086435104 B/op	     306 allocs/op
BenchmarkTwoSum/OnePassV3-16          	       1	1177603900 ns/op	32086429704 B/op	     302 allocs/op
BenchmarkTwoSum/OnePassV4-16          	     990	   1376776 ns/op	 1312401 B/op	    2207 allocs/op
BenchmarkTwoSum/OnePassV5-16          	    1248	   1092504 ns/op	 1312579 B/op	    2210 allocs/op
PASS
ok  	leetcode-fighter-go/array_and_hasing/two_integers_sum	56.388s


This is so damn weird and contra to pretty much the theory. The only possible explanation is the overhead time for memory
allocation for the map. As you can see, in the case when N = 10,000,000, the performance of V3 and V4 is widly different. 
V4 is **855 times faster** than V3. Worse, Bruteforce function is **twice** time faster than V3 as well! What the hell is this? It is shocking to know that one slight modification change the 
game:
>`dict_input := make(map[int]int, len(input.ArrayInt))` to `dict_input := make(map[int]int)`

In the other hand, compare Brute Force and OnePass V4 with N = 10,000,000, it is easily to regconize that OnePass V4 is 365 times
faster than BruteForce. This is, however, not true if N < 10,000. You can check the benchmark in the .txt file (attached in the project repo)

The takeaway lession could be: **"Becareful when preallocation!"**

#### Performance in compared with Neetcode solution

It's undisputable that my optimal strategy is somehow less performance than Neetcode. Which is pretty weird. After various
twist and change I still don't understand why. However, I figured out that the test number #5 is always less performant than
#6 (no matter what algorithm I used). I believe this is the problem in the test suits itself. 

>goos: windows
goarch: amd64
pkg: leetcode-fighter-go/array_and_hasing/two_integers_sum
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkTwoSum/Bruteforce-16         	       2	 674462050 ns/op	    1604 B/op	     100 allocs/op
BenchmarkTwoSum/OnePassV1-16          	       1	18880417400 ns/op	68694848 B/op	   27573 allocs/op
BenchmarkTwoSum/OnePassV2-16          	       1	14448088200 ns/op	32086432448 B/op	     304 allocs/op
BenchmarkTwoSum/OnePassV3-16          	       1	1132014100 ns/op	32086432448 B/op	     304 allocs/op
**BenchmarkTwoSum/OnePassV4-16          	    1125	   1371105 ns/op	 1201097 B/op	    2176 allocs/op**
BenchmarkTwoSum/OnePassV5-16          	    1200	    974888 ns/op	 1201155 B/op	    2177 allocs/op
BenchmarkTwoSum/OnePassV5#01-16       	    1340	   1028493 ns/op	 1200973 B/op	    2174 allocs/op
BenchmarkTwoSum/OnePassV5#02-16       	    1245	    978783 ns/op	 1201055 B/op	    2175 allocs/op
BenchmarkTwoSum/OnePassV5#03-16       	    1311	    943114 ns/op	 1201110 B/op	    2177 allocs/op
**BenchmarkTwoSum/OnePassV4#01-16       	    1257	    930354 ns/op	 1200967 B/op	    2174 allocs/op**
PASS
ok  	leetcode-fighter-go/array_and_hasing/two_integers_sum	62.693s


### Conclusion

Preallocation can be dangerous! Memory allocation itself is very expensive!


