package task_347

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	nums     []int
	k        int
	expected []int
}

func TestTopKFrequent(t *testing.T) {
	testCases := []TestCase{
		{"Два наиболее встречающихся в отсортированном массиве", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"Два наиболее встречающихся в неотсортированном массиве", []int{3, 2, 2, 1, 1, 1}, 2, []int{1, 2}},
		{"Единственный элемент", []int{1}, 1, []int{1}},
		{"Макс кол-во встречающихся", []int{1, 1, 1, 2, 2, 3}, 3, []int{1, 2, 3}},
		{"Отрицательные числа", []int{-1, -1, -1, 2, 2, 3}, 2, []int{-1, 2}},
		{"Отрицательные числа 2", []int{4, 1, -1, 2, -1, 2, 3}, 2, []int{-1, 2}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.nums)

			actual := topKFrequent(testCase.nums, testCase.k)
			assert.ElementsMatch(t, testCase.expected, actual)
		})
	}
}
