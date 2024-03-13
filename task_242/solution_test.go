package task_242

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	s        string
	t        string
	expected bool
}

func TestIsAnagram(t *testing.T) {
	testCases := []TestCase{
		{"Одинаковая длина строки и одинаковые символы", "anagram", "nagaram", true},
		{"Одинаковая длина строки, но разные символы", "rat", "car", false},
		{"Разная длина строк и разные символы", "a", "ab", false},
		{"Разная длина строк, но одинаковые символы", "a", "aa", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := isAnagram(testCase.s, testCase.t)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
