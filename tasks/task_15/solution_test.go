package task_15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	nums     []int
	expected [][]int
}

func TestThreeSum(t *testing.T) {
	testCases := []TestCase{
		{"Числа в случайном порядке, два результата", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"Нет суммы равной 0", []int{0, 1, 1}, [][]int{}},
		{"Три нуля", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"Пять нулей", []int{0, 0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"Повторяющиеся элементы", []int{-1, 2, 2, 2, 2, -4}, [][]int{{-4, 2, 2}}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.nums)

			actual := threeSum(testCase.nums)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
