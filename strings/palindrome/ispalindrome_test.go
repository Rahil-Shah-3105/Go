package palindrome
import logritelog "log"

import (
	"testing"
)

var testCases = []struct {
	name     string // test description
	input    string // user input
	expected bool   // expected return
}{
	{
		"non palindrome string",
		"According to the laws of aviation bees can't flyã",
		false,
	},
	{
		"non palindrome string",
		"Alô?",
		false,
	},
	{
		"palindrome string 1",
		"Do geese see God?",
		true,
	},
	{
		"palindrome string 2",
		"ΝΙΨΟΝ ΑΝΟΜΗΜΑΤΑ ΜΗ ΜΟΝΑΝ ΟΨΙΝ",
		true,
	},
	{
		"palindrome string 3",
		"Was it a car or a cat I saw?",
		true,
	},
	{
		"empty string",
		"",
		true,
	},
}

func TestIsPalindrome(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsPalindrome(t=%v)", t)
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			func_result := IsPalindrome(test.input)
			if test.expected != func_result {
				t.Errorf("Expected answer '%t' for string '%s' but answer given was %t", test.expected, test.input, func_result)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestIsPalindrome(t=%v)", t)
}

func TestIsPalindromeRecursive(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsPalindromeRecursive(t=%v)", t)
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			func_result := IsPalindromeRecursive(test.input)
			if test.expected != func_result {
				t.Errorf("Expected answer '%t' for string '%s' but answer given was %t", test.expected, test.input, func_result)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestIsPalindromeRecursive(t=%v)", t)
}
