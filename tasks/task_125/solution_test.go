package task_125

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	s        string
	expected bool
}

/*
Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/
func TestIsPalindrome(t *testing.T) {
	testCases := []TestCase{
		{"Строка палиндром со знаками препинания", "A man, a plan, a canal: Panama", true},
		{"Строка не палиндром", "race a car", false},
		{"Пустая строка", "", true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := isPalindrome(testCase.s)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
