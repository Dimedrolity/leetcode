package task_238

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	nums     []int
	expected []int
}

func TestProductExceptSelf(t *testing.T) {
	testCases := []TestCase{
		{"Массив из двух элементов", []int{2, 3}, []int{3, 2}},
		{"Массив из трех элементов", []int{2, 3, 5}, []int{15, 10, 6}},
		{"Массив положительных чисел", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"Массив с отрицательными числами и одним нулем", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{"Массив с двумя нулями", []int{0, 1, 1, 2, 0}, []int{0, 0, 0, 0, 0}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.nums)

			actual := productExceptSelf(testCase.nums)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
