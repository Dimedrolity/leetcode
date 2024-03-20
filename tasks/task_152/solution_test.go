package task_152

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	nums     []int
	expected int
}

func TestMaxProduct(t *testing.T) {
	testCases := []TestCase{
		{"Массив с одним отрицательным числом", []int{2, 3, -2, 4}, 6},
		{"Массив с четным количеством отрицательных чисел", []int{-2, -2, -2, -2}, 16},
		{"Массив с четным количеством отрицательных чисел и нулем", []int{-2, 0, -1}, 0},
		{"Массив с двумя отрицательными числами, которые дают макс результат", []int{-2, -2, 0, -1}, 4},
		{"Массив с двумя отрицательными числами и двумя нулями", []int{-2, 0, -2, 0, -1}, 0},
		{"Массив с отрицательными числами и нулем, макс результат - это один из элементов массива", []int{-3, 0, 1, -2}, 1},
		{"Массив с отрицательными числами и нулем, макс результат - подмассив не с начала массива", []int{-1, -2, -3, 0}, 6},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.nums)

			actual := maxProduct(testCase.nums)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
