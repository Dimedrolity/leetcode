package task_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	s        string
	expected int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []TestCase{
		{"Строка с повторяющимися элементами и уникальными в начале", "abcabcbb", 3},
		{"Строка из одного уникального элемента", "bbbbb", 1},
		{"Два подряд одинаковых", "pwwkew", 3},
		{"Один элемент", "й", 1},
		{"Пустая строка", "", 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := lengthOfLongestSubstring(testCase.s)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
