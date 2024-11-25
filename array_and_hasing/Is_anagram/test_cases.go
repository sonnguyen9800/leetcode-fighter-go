package is_anagram

var testCases = []IsAnagramTestCase{
	// Basic anagrams
	{"listen", "silent", true},
	{"evil", "vile", true},
	{"rat", "tar", true},
	{"angel", "glean", true},
	{"binary", "brainy", true},
	{"debitcard", "badcredit", true},
	{"dormitory", "dirtyroom", true},
	{"theeyes", "theysee", true},
	{"admirer", "married", true},
	{"funeral", "realfun", true},

	// Not anagrams
	{"hello", "world", false},
	{"cat", "rat", false},
	{"good", "dog", false},
	{"night", "thingy", false},
	{"listen", "listing", false},
	{"test", "testing", false},
	{"node", "donee", false},
	{"tracer", "reactor", false},
	{"statue", "status", false},
	{"flow", "wolfish", false},

	// Mixed case anagrams (case-sensitive)
	{"Listen", "Silent", true},
	{"Evil", "Vile", true},
	{"Tea", "Eat", true},
	{"Part", "Trap", true},
	{"Cat", "Act", true},
	{"Stressed", "Desserts", true},
	{"Elbow", "Below", true},
	{"Dusty", "Study", true},
	{"Angel", "Angle", true},
	{"Brag", "Grab", true},

	// Case-insensitive anagrams (lowercase equivalent anagrams)
	{"Astronomer", "Moonstarrer", false},
	{"Conversation", "Voicesranton", true},
	{"TheMorseCode", "Herecomethedots", false},
	{"ElevenplusTwo", "TwelveplusOne", true},
	{"Dormitory", "Dirtyroom", true},

	// Edge cases (empty strings and single characters)
	{"", "", true},         // Both strings empty
	{"a", "a", true},       // Single character, same
	{"a", "b", false},      // Single character, different
	{"ab", "a", false},     // Different lengths
	{"", "a", false},       // One string empty
	{"ab", "ba", true},     // Simple 2-char anagram
	{"abcd", "dcba", true}, // Simple 4-char anagram
	{"abcd", "abdc", true}, // Simple shuffle

	// Repeating characters
	{"aabbcc", "abcabc", true},
	{"aaabb", "aabba", true},
	{"abcabc", "cbaabc", true},
	{"abcdabcd", "abcddcba", true},
	{"zzxxcc", "xzxczc", true},
	{"aaa", "aaaa", false},          // Different counts
	{"aaa", "bbb", false},           // Completely different
	{"aaaaaaaa", "aaabaaaa", false}, // Almost same

	// Unicode anagrams
	{"你好", "好你", true},         // Simple Chinese anagram
	{"ありがとう", "とうありがと", false}, // Japanese character shuffle
	{"привет", "тпевир", true}, // Cyrillic letters anagram
	{"مرحبا", "باحرم", true},   // Arabic letters anagram
	{"γειά", "άγει", true},     // Greek letters anagram
	{"你好", "你你", false},        // Non-anagram Chinese characters
	{"ありがとう", "ありがとうう", false}, // Non-anagram Japanese
	{"مرحبا", "مرحبه", false},  // Almost similar Arabic
	{"γειά", "γειάγει", false}, // Additional Greek characters

	//// Long strings with identical characters
	{string(make([]byte, 10000)), string(make([]byte, 10000)), true}, // All-zero long string
	{string(make([]byte, 9999)), string(make([]byte, 10000)), false}, // Different lengths

	// Palindromic anagrams
	{"civic", "vicci", true},
	{"level", "evell", true},
	{"madam", "amdam", true},
	{"radar", "ardra", true},
	{"refer", "efrer", true},

	// Long shuffled anagrams
	{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba", true},
	{"thequickbrownfoxjumpsoverthelazydog", "godlazytheoverjumpsfoxbrownquickthe", true},
	{"racecar", "caracer", true},

	// Special characters
	{"123", "321", true},
	{"@!#$", "$#@!", true},
	{"123abc", "abc321", true},
	{"hello!", "oellh!", true},
	{"abc#", "#cba", true},
	{"abc!", "a!bc", true},
	{"123@", "@321", true},

	// Different character sets with similar lengths
	{"abcde", "vwxyz", false},
	{"12345", "54321", true},
	{"@#$%^", "%^$#@", true},
	{"@@@@@", "####", false},

	// Randomized pairs
	{"abcdefgh", "hgfedcba", true},

	{"aabbccdd", "ddccbbaa", true},
	{"mississippi", "ipissmisspi", true},
	{"california", "niforcalia", true},
	{"newyork", "yorknewy", false},
	{"texas", "saext", true},

	// Extreme character mix
	{"AaBbCcDdEeFfGg", "GgFfEeDdCcBbAa", true},
	{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba", true},
	{"abcdabcdabcdabcd", "dcbaabcdabcdabcd", true},
	{"xyzxyzxyzxyzxyz", "yxzxyzyzxyzyxyz", false},

	// Extreme character mix with mixed cases and shuffled order
	{"AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz",
		"ZzYyXxWwVvUuTtSsRrQqPpOoNnMmLlKkJjIiHhGgFfEeDdCcBbAa",
		true},

	// Reversed alphabet
	{"abcdefghijklmnopqrstuvwxyz",
		"zyxwvutsrqponmlkjihgfedcba",
		true},

	// Repeated pattern of a small set of characters
	{"abcdabcdabcdabcdabcdabcd",
		"dcbaabcdabcdabcdabcdabcd",
		true},

	// Character set with a mix of digits and letters
	{"1234567890abcdefghijklmnopqrstuvwxyz",
		"abcdefghijklmnopqrstuvwxyz1234567890",
		true},

	// Single character repeated multiple times
	{"aaaaaaa",
		"aaaaaaa",
		true},

	// Edge case with one character mismatch
	{"abcd",
		"abce",
		false},

	// Large strings with many repeated characters
	{"aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz",
		"zzxxwwvvuuttsrrqqppoonnmmllkkjjiihhggffeeddccbbaa",
		false},

	// Long strings with a mixture of case and order
	{"aAaaAaBbCcCcBbAaAa",
		"BbAaAaaBCAaCcBa",
		false},

	// Strings with one additional character in one string
	{"xyzxyzxyzxyzxyzxyz",
		"yzxyzxyzxyzxyzxyza",
		false},

	// Extreme case with empty strings
	{"",
		"",
		true},
}
