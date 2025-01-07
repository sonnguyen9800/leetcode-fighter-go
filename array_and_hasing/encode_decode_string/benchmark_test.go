package encode_decode_string

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

const (
	TEST_CASES_NUMBER = 1
	MIN_STRING_LENGTH = 0
	MAX_STRING_LENGTH = 20
	MIN_ARRAY_SIZE    = 0
	MAX_ARRAY_SIZE    = 100
	CHARSET           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 !@#$%^&*()-_=+[]{}|;:'\",.<>?/"
)

func generateRandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = CHARSET[rand.Intn(len(CHARSET))]
	}
	return string(result)
}

func generateTestCase() []string {
	size := rand.Intn(MAX_ARRAY_SIZE-MIN_ARRAY_SIZE+1) + MIN_ARRAY_SIZE
	var strs []string

	for i := 0; i < size; i++ {
		length := rand.Intn(MAX_STRING_LENGTH - MIN_STRING_LENGTH + 1)
		strs = append(strs, generateRandomString(length))
	}

	return strs
}

func generateFuzzyTestCases(numCases int) [][]string {
	var testCases [][]string
	for i := 0; i < numCases; i++ {
		testCases = append(testCases, generateTestCase())
	}
	return testCases
}

func BenchmarkEncodeDecode(b *testing.B) {

	testCases := generateFuzzyTestCases(TEST_CASES_NUMBER)

	b.ResetTimer()
	b.Run("EncodeDecode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tc := range testCases {
				if len(tc) == 0 {
					continue
				}
				encoded := encode_string(tc)
				decoded := decode_string(encoded)

				if !reflect.DeepEqual(tc, decoded) {
					b.Errorf("Failed:\n")

					for _, s := range tc {
						for _, char := range s {
							fmt.Printf("Character: %d ", char) // Print each character as int
						}
					}
					b.Logf("Original: |%s| %d element", strings.Join(tc, " |"), len(tc))
					b.Logf("Encoded: |%v|", encoded)
					b.Logf("Decoded: |%s| %d element", strings.Join(decoded, " |"), len(tc))

				}
			}
		}
	})
}

func BenchmarkEncodeDecode2(b *testing.B) {
	testCases := []string{}

	b.ResetTimer()
	b.Run("EncodeDecode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			encoded := encode_string(testCases)
			decoded := decode_string(encoded)

			if !reflect.DeepEqual(testCases, decoded) {
				b.Errorf("Failed:\n")
				b.Logf("Original: |%s| %d element", strings.Join(testCases, " |"), len(testCases))
				b.Logf("Encoded: |%v|", encoded)
				b.Logf("Decoded: |%s", strings.Join(decoded, " |"))
			}
		}
	})
}
