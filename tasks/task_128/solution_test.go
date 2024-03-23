package task_128

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	nums     []int
	expected int
}

func TestLongestConsecutive(t *testing.T) {
	testCases := []TestCase{
		{"Три последовательности (от 1 до 4, 100 и 200)", []int{100, 4, 200, 1, 3, 2}, 4},
		{"Одна последовательность из элементов в случайном порядке", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"Ноль элементов", []int{}, 0},
		{"Один элемент", []int{0}, 1},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.nums)

			actual := longestConsecutive(testCase.nums)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
