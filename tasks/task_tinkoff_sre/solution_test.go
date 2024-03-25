package task_tinkoff_sre

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	s        string
	expected string
}

func TestContainsDuplicate(t *testing.T) {
	testCases := []TestCase{
		{"Без дублей", "ABC", "ABC"},
		{"Один дубль", "AA", "A2"},
		{"Дубль и символ", "AAB", "A2B"},
		{"Два дубля", "AAAABBBC", "A4B3C"},
		{"Четыре дубля", "AAAABBBCAAA", "A4B3CA3"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := encode(testCase.s)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
