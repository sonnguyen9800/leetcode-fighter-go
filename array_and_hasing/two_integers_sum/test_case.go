package main

import (
	"math/rand"
)

type TwoIntegersSumTestCase struct {
	Id       int
	Input    TwoIntergersInput
	Expected [2]int
}
type TwoIntergersInput struct {
	ArrayInt []int
	Target   int
}

func ConvertTestCase(in TwoIntegersSumTestCase) struct {
	input TwoIntergersInput
	want  [2]int
} {

	return struct {
		input TwoIntergersInput
		want  [2]int
	}{
		input: in.Input,
		want:  in.Expected,
	}

}

var testCases = []TwoIntegersSumTestCase{
	{Input: TwoIntergersInput{ArrayInt: []int{2, 7, 11, 15}, Target: 9}, Expected: [2]int{0, 1}},
	{Input: TwoIntergersInput{ArrayInt: []int{3, 2, 4}, Target: 6}, Expected: [2]int{1, 2}},
	{Input: TwoIntergersInput{ArrayInt: []int{3, 3}, Target: 6}, Expected: [2]int{0, 1}},
	{Input: TwoIntergersInput{ArrayInt: []int{5, 25, 75}, Target: 100}, Expected: [2]int{1, 2}},
	{Input: TwoIntergersInput{ArrayInt: []int{0, 4, 3, 0}, Target: 0}, Expected: [2]int{0, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{-3, 4, 3, 90}, Target: 0}, Expected: [2]int{0, 2}},
	{Input: TwoIntergersInput{ArrayInt: []int{-1, -2, -3, -4, -5}, Target: -8}, Expected: [2]int{2, 4}},
	{Input: TwoIntergersInput{ArrayInt: []int{1, 6, 7, 8}, Target: 15}, Expected: [2]int{2, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{1, 5, 1, 5}, Target: 10}, Expected: [2]int{1, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{2, 7, 2, 8}, Target: 10}, Expected: [2]int{0, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{4, 9, 11, 15}, Target: 20}, Expected: [2]int{1, 2}},
	{Input: TwoIntergersInput{ArrayInt: []int{5, 3, 10, 6, 8}, Target: 14}, Expected: [2]int{3, 4}},
	{Input: TwoIntergersInput{ArrayInt: []int{20, 5, 25, 35}, Target: 55}, Expected: [2]int{0, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{-3, -2, 7, 5, -7}, Target: 0}, Expected: [2]int{2, 4}},
	{Input: TwoIntergersInput{ArrayInt: []int{6, 10, 15, 3}, Target: 9}, Expected: [2]int{0, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{8, 15, 7, 3}, Target: 11}, Expected: [2]int{0, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{1, 3, 4, 2}, Target: 6}, Expected: [2]int{2, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{1, 3, 5, 2}, Target: 7}, Expected: [2]int{2, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{0, 1, 2, 0}, Target: 0}, Expected: [2]int{0, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{-5, -10, -3, 5}, Target: 0}, Expected: [2]int{0, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{15, 5, 7, 10}, Target: 12}, Expected: [2]int{1, 2}},
	{Input: TwoIntergersInput{ArrayInt: []int{9, 3, 1, 6}, Target: 10}, Expected: [2]int{0, 2}},
	{Input: TwoIntergersInput{ArrayInt: []int{2, 11, 7, 13}, Target: 9}, Expected: [2]int{0, 2}},
	{Input: TwoIntergersInput{ArrayInt: []int{8, 5, 3, 2}, Target: 8}, Expected: [2]int{1, 2}},
	{Input: TwoIntergersInput{ArrayInt: []int{-10, 10, -5, 15}, Target: 0}, Expected: [2]int{0, 1}},
	{Input: TwoIntergersInput{ArrayInt: []int{6, 14, 10, 4}, Target: 10}, Expected: [2]int{0, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{-2, 7, -3, 10}, Target: 5}, Expected: [2]int{0, 1}},
	{Input: TwoIntergersInput{ArrayInt: []int{6, -2, 3, 9}, Target: 7}, Expected: [2]int{1, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{8, -8, 7, 1}, Target: 0}, Expected: [2]int{0, 1}},
	{Input: TwoIntergersInput{ArrayInt: []int{7, 6, 2, 8}, Target: 14}, Expected: [2]int{1, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{10, 20, 30, 40}, Target: 50}, Expected: [2]int{0, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{0, 5, 5, 0}, Target: 10}, Expected: [2]int{1, 2}},
	{Input: TwoIntergersInput{ArrayInt: []int{1, 4, 4, 1}, Target: 8}, Expected: [2]int{2, 1}},
	{Input: TwoIntergersInput{ArrayInt: []int{2, 5, 5, 11}, Target: 10}, Expected: [2]int{1, 2}},
	{Input: TwoIntergersInput{ArrayInt: []int{5, 3, 4, 3}, Target: 6}, Expected: [2]int{1, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{9, 1, 8, 3}, Target: 10}, Expected: [2]int{0, 1}},
	{Input: TwoIntergersInput{ArrayInt: []int{1, 6, 3, 7}, Target: 10}, Expected: [2]int{2, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{-1, 2, -3, 4}, Target: 1}, Expected: [2]int{0, 1}},
	{Input: TwoIntergersInput{ArrayInt: []int{4, 5, 6, 7}, Target: 13}, Expected: [2]int{2, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{11, 22, 33, 44}, Target: 55}, Expected: [2]int{0, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{1, 2, 2, 3}, Target: 4}, Expected: [2]int{0, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{7, 9, 3, 8}, Target: 16}, Expected: [2]int{1, 0}},
	{Input: TwoIntergersInput{ArrayInt: []int{6, 10, 4, 2}, Target: 12}, Expected: [2]int{1, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{1, 8, 7, 2}, Target: 9}, Expected: [2]int{1, 0}},
	{Input: TwoIntergersInput{ArrayInt: []int{4, 6, 10, 1}, Target: 11}, Expected: [2]int{2, 3}},
	{Input: TwoIntergersInput{ArrayInt: []int{7, 3, 8, 5}, Target: 15}, Expected: [2]int{0, 2}},
	{Input: TwoIntergersInput{ArrayInt: []int{4, 3, 1, 5}, Target: 8}, Expected: [2]int{1, 3}},
	// Add more combinations using similar logic to cover edge and normal cases.
}

// GenerateFuzzyTestCases generates `n` test cases, ensuring exactly one valid pair exists per input.
func GenerateFuzzyTestCases(n int,
	arraySizeRange [2]int,
	valueRange [2]int) []TwoIntegersSumTestCase {
	var outputTestCase []TwoIntegersSumTestCase

	for i := 0; i < n; i++ {
		// Generate a random array size
		arraySize := rand.Intn(arraySizeRange[1]-arraySizeRange[0]+1) + arraySizeRange[0]
		array := make([]int, 0)
		usedNumbers := make(map[int]bool) // Tracks already used numbers

		// Fill the array with unique random values
		for len(array) < arraySize {
			num := rand.Intn(valueRange[1]-valueRange[0]+1) + valueRange[0]
			if !usedNumbers[num] {
				array = append(array, num)
				usedNumbers[num] = true
			}
		}

		// Select two distinct random indices
		idx1 := rand.Intn(arraySize)
		idx2 := rand.Intn(arraySize)

		for idx2 == idx1 {
			idx1 = rand.Intn(arraySize)
			idx2 = rand.Intn(arraySize)
		}

		// Calculate the target value based on the selected indices
		target := array[idx1] + array[idx2]

		// Add the test case
		outputTestCase = append(outputTestCase, TwoIntegersSumTestCase{
			Id: i + 1,
			Input: TwoIntergersInput{
				ArrayInt: array,
				Target:   target,
			},
			Expected: [2]int{idx1, idx2},
		})
	}

	return outputTestCase
}
