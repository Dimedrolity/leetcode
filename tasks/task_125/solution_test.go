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

func TestIsPalindrome(t *testing.T) {
	testCases := []TestCase{
		{"Строка палиндром со знаками препинания", "A man, a plan, a canal: Panama", true},
		{"Строка не палиндром", "race a car", false},
		{"Пустая строка", "", true},
		{"Строка не палиндром из двух элементов", "0P", false},
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
