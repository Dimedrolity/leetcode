package task_167

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	testCases := []TestCase{
		{"Первые 2 элемента", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"Первый и последний элемент", []int{2, 3, 4}, 6, []int{1, 3}},
		{"Отрицательное число и ноль", []int{-1, 0}, -1, []int{1, 2}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.nums)

			actual := twoSum(testCase.nums, testCase.target)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
