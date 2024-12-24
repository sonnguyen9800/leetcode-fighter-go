# Top K Frequent Elements

## Problem Statement

Given an integer array `nums` and an integer `k`, return the `k` most frequent elements within the array.

The test cases are generated such that the answer is always **unique**.

You may return the output in **any order**.

*Example 1:*
```
Input: nums = [1,2,2,3,3,3], k = 2
Output: [2,3]
```
*Example 2:*

```
Input: nums = [7,7], k = 1
Output: [7]
```

### Constraints:

```
1 <= nums.length <= 10^4.
-1000 <= nums[i] <= 1000
1 <= k <= number of distinct elements in nums.
```

## A short analysis

The problem super interesting. To be honest, I couldn't figure out how to do until I read the second hint from Neetcode, that suggest grouping number *based on their frequency*

The first solution I had in my mind could be:

- Create a hashmap with key to be the number, and value to be the frequency of that number (Hashmap #1)
- From Hashmap #1, create a hashmap of frequency, with key to be the frequency and value to be an array of number that had that same frequency

A quick representation could be something like this

```
nums = [1 1 1 33 3 4 5 5]

hashmap_1 = {
	1 : 3 // value : key
	33 : 1
	3 : 1
	4: 1
	5 : 2
}

hashmap_2 = {
	3 : [1]
	1 : [33, 3, 4]
	2 : [5]
}
```

My naive thinking at that time were to reference to the key with highest frequency in Hashmap#2, take that one out for the `output_array`. But this couldn't be correct and guarantee the output to be enough (`k` can be greater than size of number inside the array of `highest frequency`). So my next in-head solution would be to keep track of two number: `max_value` and `min_value`, which represent the number with highest frequency and lowest frequency already stayed inside the `output_array`, if the next iteration in `Hashmap#2` has higher frequency, then popout the `min_value` and so on for the lowest frequency. **However**, this approach leads to a big issue: How to find the next `lowest frequency?` Is this still somehow related to sorting? If inside the loop I still need to find the next `lowest_frequency` and `highest_frequency`, this would make the `BigO` to be `O(n^2)`, which is **unacceptable**

Then, based on previous solving experience with Leetcode, I am now more familiar to the concept of **Scenario Abusing** (aka. Abusing the given constrants). Look at the constrant again and you will immediately regconite that they are made *to be abused! To be cheated at!* How easy, memory saving to make an array of `10,000` index? So cheap!

So the solution now is to looping through an array of `10,000` index, each represent the `possible frequency` of all number inside. I would create one below based on previous example
```
0	[]	
1	[33 3 4]	
2	[1]	
3	[1]
4	[]	
5	[]
... 
9,999 []
```

Of course! We don't even need to *create the actual array of 10,000 index*, but rather what we should do is simply iterate from number `9,999` to `0` by iterator `i`! Every time we loop, we append `hashmap_2[i]` to `output_array` based on the wanted `k` and current size of `output_array` itself (to avoid appending more number than required!)


### Solution:

#### First Attempt & Benchmark

```
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
			var amount int
			if k > current_length {
				amount = k - current_length
			} else {
				amount = current_length - k
			}
			output_array = append(output_array, freqMapByCount[i][0:amount]...)
			break
		} else {
			output_array = append(output_array, freqMapByCount[i]...)
		}
	}
	return output_array

}
```


But The benchmark is surprisingly disapponting!

```
goos: windows
goarch: amd64
pkg: leetcode-fighter-go/array_and_hasing/top-k-freq-element
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkTopKFreq/Custom_Solution-16						  13	  89550054 ns/op	  546803 B/op	   15555 allocs/op
BenchmarkTopKFreq/Sorting_-_Neetocode_Solution-16      	      31	  39695129 ns/op	  600144 B/op	   16271 allocs/op
BenchmarkTopKFreq/Heap_(NeetcodeSolution)-16           	      30	  39803170 ns/op	  673857 B/op	   18782 allocs/op
BenchmarkTopKFreq/BucketSort(NeetcodeSolution)-16      	      22	  45516950 ns/op	97804577 B/op	   16555 allocs/op
BenchmarkTopKFreq/Using_3rd_Party(Neetcode_Solution)-16       31	  37753023 ns/op	  673857 B/op	   18782 allocs/op
PASS
ok  	leetcode-fighter-go/array_and_hasing/top-k-freq-element	7.495s
```

The `CustomSolution` has highest number of `time spent`. The only good point is it's the most memory efficient! A more detailed looks reveal some disadvantages my solution had
1. Too many cache: Yes, I am not entirely familiar with `go-lang`, so I put myself into so many small, un-optimized pieces of code.
2. I don't need to fix the iterator from `9999`, I could catch the highest frequency first and loop from that frequency to 0!
3. The algorithm to `append` number inside the `output_array` is a bit too ... much! Too many `if-else` here (to prevent redundant number not to be appened). The most optimal way would be: *Just appending and return `output_array` when the size reach `k`!*. Yes! Why I hadn't think of that?
   
So the below solution is a more optimal one!

### Second Attempt: Optimal method

Less Code! More performant!
```
func TopKFreqElement_Optimal(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num] += 1
	}

	// Creating a map of frequency of each frequency
	freqMapByCount := make(map[int][]int)
	highest_frequency := 0

	for key, freqValue := range freqMap {
		freqMapByCount[freqValue] = append(freqMapByCount[freqValue], key)
		if freqValue > highest_frequency {
			highest_frequency = freqValue
		}
	}

	output_array := []int{}

	for i := highest_frequency; i >= 0; i-- {
		for _, num := range freqMapByCount[i] {
			output_array = append(output_array, num)
			if len(output_array) == k {
				return output_array
			}
		}

	}
	return output_array
}

```

The benchmark is dramatically different!

```
goos: windows
goarch: amd64
pkg: leetcode-fighter-go/array_and_hasing/top-k-freq-element
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkTopKFreq/Custom_Solution-16         	      		  15	  82004513 ns/op	  551282 B/op	   15965 allocs/op
BenchmarkTopKFreq/Custom_Solution_-_Optimal-16         	      34	  33873232 ns/op	  551281 B/op	   15965 allocs/op
BenchmarkTopKFreq/Sorting_-_Neetocode_Solution-16      	      36	  33196022 ns/op	  607152 B/op	   16613 allocs/op
BenchmarkTopKFreq/Heap_(NeetcodeSolution)-16           	      34	  33254506 ns/op	  683168 B/op	   19206 allocs/op
BenchmarkTopKFreq/BucketSort(NeetcodeSolution)-16      	      27	  39489996 ns/op	97038882 B/op	   16965 allocs/op
BenchmarkTopKFreq/Using_3rd_Party(Neetcode_Solution)-16       37	  33143054 ns/op	  683168 B/op	   19206 allocs/op
PASS
ok  	leetcode-fighter-go/array_and_hasing/top-k-freq-element	7.450s
```

>BenchmarkTopKFreq/Custom_Solution_-_Optimal-16         	      34	  **33873232** ns/op	  551281 B/op	   15965 allocs/op
>BenchmarkTopKFreq/BucketSort(NeetcodeSolution)-16      	      27	  39489996 ns/op		97038882 B/op	   16965 allocs/op
>BenchmarkTopKFreq/Using_3rd_Party(Neetcode_Solution)-16       	  37	  33143054 ns/op	  	  683168 B/op	   19206 allocs/op

My `optimal solution` is simply the best overall! You can argue that the solution using `sorting` is more efficient BUT it's in `O(nlogn)!` 

The `neetcode solution using 3rd party` (I couldn't find the reference to it on the internet. Not sure where I grabbed it!) Looked to be somehow similar, however, it is using third party library. They could likely do something optimized by writing in lower level code (maybe?)

The newer version also not require any fixed number (`9999`)! This would theoretically support for maximum size possible. Now, let do the test with `N > 10,000`!


### N = 100,000

Let's see what happened when testing with N = 100,000, what I hoped for is that the `Sorting Solution` would be fallen back behind the `optimal` and `bucket sort` approach.

```
goos: windows
goarch: amd64
pkg: leetcode-fighter-go/array_and_hasing/top-k-freq-element
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkTopKFreq/Custom_Solution_-_Optimal-16         	       1	2170082000 ns/op	  572320 B/op	   17915 allocs/op
BenchmarkTopKFreq/Sorting_-_Neetocode_Solution-16      	       1	2181094500 ns/op	  641632 B/op	   18275 allocs/op
BenchmarkTopKFreq/Heap_(NeetcodeSolution)-16           	       1	2161100800 ns/op	  729472 B/op	   21260 allocs/op
BenchmarkTopKFreq/BucketSort(NeetcodeSolution)-16      	       1	2814506200 ns/op	4760282736 B/op	   18918 allocs/op
BenchmarkTopKFreq/Using_3rd_Party(Neetcode_Solution)-16        1	2146205400 ns/op	  729472 B/op	   21258 allocs/op
PASS
ok  	leetcode-fighter-go/array_and_hasing/top-k-freq-element	15.935s
```

Is this a sick joke? `Sorting approach` which supposed to be `less good` somehow stay on pair with my `optimal solution` & `third-party using solution`. Why?

```
goos: windows
goarch: amd64
pkg: leetcode-fighter-go/array_and_hasing/top-k-freq-element
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkTopKFreq/Custom_Solution_-_Optimal-16         	       3	 480298533 ns/op	  557552 B/op	   16526 allocs/op
BenchmarkTopKFreq/Sorting_-_Neetocode_Solution-16      	       3	 480922167 ns/op	  616189 B/op	   17063 allocs/op
BenchmarkTopKFreq/Heap_(NeetcodeSolution)-16           	       3	 509211533 ns/op	  694936 B/op	   19768 allocs/op
BenchmarkTopKFreq/BucketSort(NeetcodeSolution)-16      	       2	 579715250 ns/op	1152279144 B/op	   17526 allocs/op
BenchmarkTopKFreq/Using_3rd_Party(Neetcode_Solution)-16        3	 500474800 ns/op	  694930 B/op	   19768 allocs/op
PASS
ok  	leetcode-fighter-go/array_and_hasing/top-k-freq-element	14.746s
```
### N = 500,000

```
goos: windows
goarch: amd64
pkg: leetcode-fighter-go/array_and_hasing/top-k-freq-element
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkTopKFreq/Custom_Solution_-_Optimal-16         	       1	11234137200 ns/op	  572848 B/op	   17960 allocs/op
BenchmarkTopKFreq/Sorting_-_Neetocode_Solution-16      	       1	11013849300 ns/op	  642216 B/op	   18308 allocs/op
BenchmarkTopKFreq/Heap_(NeetcodeSolution)-16           	       1	11140384700 ns/op	  730136 B/op	   21298 allocs/op
BenchmarkTopKFreq/BucketSort(NeetcodeSolution)-16      	       1	14497243000 ns/op	23907125424 B/op	   18963 allocs/op
BenchmarkTopKFreq/Using_3rd_Party(Neetcode_Solution)-16  	   1	11091479300 ns/op	  730104 B/op	   21298 allocs/op
PASS
ok  	leetcode-fighter-go/array_and_hasing/top-k-freq-element	83.345s
```

The problem could be solved if the N = 1,000,000,000 or more. **But my machine already run out of memory with N=1000000**.

### Conclusion
Yay! Optimal is fun, but theory is not aways reflected in practice. The different with N < 500,000 is so damn small! There is virtually nothing change if I sort everything instead! Not fun at all!


