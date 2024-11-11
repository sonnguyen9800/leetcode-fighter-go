package main

import "leetcode-fighter-go/goland-slib/test"

type TwoIntergersInput struct {
	ArrayInt []int
	Target   int
}
type TwoIntegersSumTestCase struct {
	Input    TwoIntergersInput
	Expected [2]int
}

func (testCase TwoIntegersSumTestCase) ConvertTestcases() test.GenericOutput[TwoIntergersInput, [2]int] {
	return test.GenericOutput[TwoIntergersInput, [2]int]{
		Input:    testCase.Input,
		Expected: testCase.Expected,
	}
}

func runTests() {

	testCases := []TwoIntegersSumTestCase{
		{Input: TwoIntergersInput{ArrayInt: []int{2, 7, 11, 15}, Target: 9}, Expected: [2]int{0, 1}},
		{Input: TwoIntergersInput{ArrayInt: []int{3, 2, 4}, Target: 6}, Expected: [2]int{1, 2}},
		{Input: TwoIntergersInput{ArrayInt: []int{3, 3}, Target: 6}, Expected: [2]int{0, 1}},
		{Input: TwoIntergersInput{ArrayInt: []int{1, 2, 3, 4, 5}, Target: 6}, Expected: [2]int{1, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{5, 25, 75}, Target: 100}, Expected: [2]int{1, 2}},
		{Input: TwoIntergersInput{ArrayInt: []int{0, 4, 3, 0}, Target: 0}, Expected: [2]int{0, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{-3, 4, 3, 90}, Target: 0}, Expected: [2]int{0, 2}},
		{Input: TwoIntergersInput{ArrayInt: []int{-1, -2, -3, -4, -5}, Target: -8}, Expected: [2]int{2, 4}},
		{Input: TwoIntergersInput{ArrayInt: []int{1, 6, 7, 8}, Target: 15}, Expected: [2]int{2, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{1, 5, 1, 5}, Target: 10}, Expected: [2]int{1, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{2, 7, 2, 8}, Target: 10}, Expected: [2]int{1, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{4, 9, 11, 15}, Target: 20}, Expected: [2]int{0, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{5, 3, 10, 6, 8}, Target: 14}, Expected: [2]int{2, 4}},
		{Input: TwoIntergersInput{ArrayInt: []int{20, 5, 25, 35}, Target: 55}, Expected: [2]int{2, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{-3, -2, 7, 5, -7}, Target: 0}, Expected: [2]int{0, 4}},
		{Input: TwoIntergersInput{ArrayInt: []int{6, 10, 15, 3}, Target: 9}, Expected: [2]int{0, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{8, 15, 7, 3}, Target: 11}, Expected: [2]int{2, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{1, 3, 4, 2}, Target: 6}, Expected: [2]int{2, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{1, 3, 5, 2}, Target: 7}, Expected: [2]int{1, 2}},
		{Input: TwoIntergersInput{ArrayInt: []int{0, 1, 2, 0}, Target: 0}, Expected: [2]int{0, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{3, 2, 4, 3}, Target: 6}, Expected: [2]int{1, 2}},
		{Input: TwoIntergersInput{ArrayInt: []int{-5, -10, -3, 5}, Target: 0}, Expected: [2]int{0, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{6, 8, 10, 4}, Target: 14}, Expected: [2]int{1, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{15, 5, 7, 10}, Target: 12}, Expected: [2]int{1, 2}},
		{Input: TwoIntergersInput{ArrayInt: []int{9, 3, 1, 6}, Target: 10}, Expected: [2]int{0, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{4, 9, 6, 11}, Target: 15}, Expected: [2]int{1, 2}},
		{Input: TwoIntergersInput{ArrayInt: []int{2, 11, 7, 13}, Target: 9}, Expected: [2]int{0, 2}},
		{Input: TwoIntergersInput{ArrayInt: []int{8, 5, 3, 2}, Target: 8}, Expected: [2]int{1, 2}},
		{Input: TwoIntergersInput{ArrayInt: []int{-10, 10, -5, 15}, Target: 0}, Expected: [2]int{0, 1}},
		{Input: TwoIntergersInput{ArrayInt: []int{6, 14, 10, 4}, Target: 10}, Expected: [2]int{0, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{5, 25, 10, 5}, Target: 30}, Expected: [2]int{1, 2}},
		{Input: TwoIntergersInput{ArrayInt: []int{-2, 7, -3, 10}, Target: 5}, Expected: [2]int{0, 1}},
		{Input: TwoIntergersInput{ArrayInt: []int{6, 7, 3, 9}, Target: 13}, Expected: [2]int{0, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{4, 3, 7, 8}, Target: 11}, Expected: [2]int{1, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{12, 5, 9, 3}, Target: 15}, Expected: [2]int{1, 2}},
		{Input: TwoIntergersInput{ArrayInt: []int{10, 6, 4, 6}, Target: 12}, Expected: [2]int{0, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{-6, 10, -4, 8}, Target: 2}, Expected: [2]int{1, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{2, 3, 5, 7}, Target: 8}, Expected: [2]int{1, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{4, -4, 10, -6}, Target: 4}, Expected: [2]int{0, 1}},
		{Input: TwoIntergersInput{ArrayInt: []int{1, 7, 8, 10}, Target: 15}, Expected: [2]int{2, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{6, -2, 3, 9}, Target: 7}, Expected: [2]int{1, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{8, -8, 7, 1}, Target: 0}, Expected: [2]int{0, 1}},
		{Input: TwoIntergersInput{ArrayInt: []int{5, 15, 10, 0}, Target: 15}, Expected: [2]int{1, 3}},
		{Input: TwoIntergersInput{ArrayInt: []int{7, 6, 2, 8}, Target: 14}, Expected: [2]int{1, 3}},
	}

	var newConverted = [](test.GenericOutput[TwoIntergersInput, [2]int]){}
	for _, testcase := range testCases {

		newConverted = append(newConverted, testcase.ConvertTestcases())
	}

	test.RunTestProduceIntArray[TwoIntergersInput](newConverted, TwoSum)

}

func main() {
	runTests()
}
