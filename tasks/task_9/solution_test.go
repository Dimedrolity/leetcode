package task_217_contains_duplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	x        int
	expected bool
}

func TestContainsDuplicate(t *testing.T) {
	testCases := []TestCase{
		{"Положительное число", 121, true},
		{"Отрицательное число", -121, false},
		{"Не палиндром", 10, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := isPalindrome(testCase.x)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
