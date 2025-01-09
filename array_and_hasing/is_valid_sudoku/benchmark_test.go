package is_valid_sudoku

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

const (
	MIN_FILLED_CELLS = 10
	MAX_FILLED_CELLS = 30
	BOARD_SIZE       = 9
	TEST_CASES_NUM   = 1000
)

// Generate a random Sudoku board
func generateRandomSudokuBoard() [][]string {
	rand.Seed(time.Now().UnixNano())
	board := make([][]string, BOARD_SIZE)
	for i := 0; i < BOARD_SIZE; i++ {
		board[i] = make([]string, BOARD_SIZE)
		for j := 0; j < BOARD_SIZE; j++ {
			board[i][j] = "."
		}
	}
	numCells := rand.Intn(MAX_FILLED_CELLS-MIN_FILLED_CELLS+1) + MIN_FILLED_CELLS
	for k := 0; k < numCells; k++ {
		row := rand.Intn(BOARD_SIZE)
		col := rand.Intn(BOARD_SIZE)
		if board[row][col] == "." {
			board[row][col] = string(rune('1' + rand.Intn(9))) // Random digit between 1 and 9
		}
	}
	return board
}

// Introduce invalidity into the board
func introduceInvalidity(board [][]string, invalidType string) {
	rand.Seed(time.Now().UnixNano())
	switch invalidType {
	case "row":
		row := rand.Intn(BOARD_SIZE)
		val := string(rune('1' + rand.Intn(9)))
		board[row][rand.Intn(BOARD_SIZE)] = val
		board[row][rand.Intn(BOARD_SIZE)] = val
	case "column":
		col := rand.Intn(BOARD_SIZE)
		val := string(rune('1' + rand.Intn(9)))
		board[rand.Intn(BOARD_SIZE)][col] = val
		board[rand.Intn(BOARD_SIZE)][col] = val
	case "subgrid":
		startRow := (rand.Intn(3) * 3)
		startCol := (rand.Intn(3) * 3)
		val := string(rune('1' + rand.Intn(9)))
		board[startRow+rand.Intn(3)][startCol+rand.Intn(3)] = val
		board[startRow+rand.Intn(3)][startCol+rand.Intn(3)] = val
	}
}

// Generate multiple test cases
func generateFuzzyTestCases(numCases int) [][][]string {
	var testCases [][][]string
	for i := 0; i < numCases; i++ {
		board := generateRandomSudokuBoard()
		if rand.Intn(2) == 0 { // 50% chance to introduce invalidity
			invalidTypes := []string{"row", "column", "subgrid"}
			introduceInvalidity(board, invalidTypes[rand.Intn(len(invalidTypes))])
		}
		testCases = append(testCases, board)
	}
	return testCases
}

// Benchmark the solution
func BenchmarkIsValidSudoku(b *testing.B) {
	testCases := generateFuzzyTestCases(TEST_CASES_NUM)
	results := make([]bool, TEST_CASES_NUM)

	// Precompute expected results using a reference solution
	for i, testCase := range testCases {
		results[i] = isValidSudoku_neetcode(testCase)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i, testCase := range testCases {
			actual := is_valid_sudoku(testCase)
			if actual != results[i] {
				b.Errorf("Test case %d failed.\nExpected: %v, Got: %v\n", i+1, results[i], actual)
				b.Logf("Board:\n%v\n", testCase)
			}
		}
	}
}
func BenchmarkIsValidSudoku_Optimal(b *testing.B) {
	testCases := generateFuzzyTestCases(TEST_CASES_NUM)
	results := make([]bool, TEST_CASES_NUM)

	// Precompute expected results using a reference solution
	for i, testCase := range testCases {
		results[i] = isValidSudoku_neetcode(testCase)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i, testCase := range testCases {
			actual := is_valid_sudoku_optimal(testCase)
			if actual != results[i] {
				b.Errorf("Test case %d failed.\nExpected: %v, Got: %v\n", i+1, results[i], actual)
				b.Logf("Board:\n%v\n", testCase)
			}
		}
	}
}

func Convert(input string) [][]string {
	// Step 1: Remove outer brackets
	input = strings.Trim(input, "[]")

	// Step 2: Split by row (separated by "] [")
	rows := strings.Split(input, "] [")
	testCases := [][]string{}
	for _, row := range rows {
		// Remove any remaining brackets and split by spaces
		row = strings.Trim(row, "[]")
		elements := strings.Fields(row)
		testCases = append(testCases, elements)
	}
	return testCases
}

func BenchmarkIsValidSudoku_singlecase(b *testing.B) {
	testCases := make([][]string, 1)
	errorCaseStr := "        " +
		"[[. . . . . 8 . . .] " +
		"[. 2 9 . . . . . .] " +
		"[. 6 . . . . . . .] " +
		"[6 . . 7 . . . 2 .] " +
		"[. . . . . . 8 . 2] " +
		"[. . . . . . . . .] " +
		"[. . . . . . . . .] " +
		"[. 9 . . . . 1 . .] " +
		"[. 7 . . . . . . .]]"

	testCases = Convert(errorCaseStr)

	// Precompute expected results using a reference solution
	results := isValidSudoku_neetcode(testCases)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		actual := is_valid_sudoku(testCases)
		if actual != results {
			b.Errorf("Test case failed.\nExpected: %v, Got: %v\n", results, actual)
			b.Logf("Board:\n%v\n", testCases)
		}
	}
}
