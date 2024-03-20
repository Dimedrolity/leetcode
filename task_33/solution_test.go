package task_33

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	nums     []int
	target   int
	expected int
}

func TestSearch(t *testing.T) {
	testCases := []TestCase{
		{"Сдвиг на 1", []int{7, 1, 2, 3, 4, 5, 6}, 1, 1},
		{"Сдвиг на 2", []int{6, 7, 1, 2, 3, 4, 5}, 1, 2},
		{"Сдвиг на 3", []int{5, 6, 7, 1, 2, 3, 4}, 1, 3},
		{"Сдвиг на 4", []int{4, 5, 6, 7, 1, 2, 3}, 1, 4},
		{"Сдвиг на 5", []int{3, 4, 5, 6, 7, 1, 2}, 1, 5},
		{"Сдвиг на 6", []int{2, 3, 4, 5, 6, 7, 1}, 1, 6},
		{"Сдвиг на 7", []int{1, 2, 3, 4, 5, 6, 7}, 1, 0},
		{"Два элемента, сдвиг на 1", []int{2, 1}, 1, 1},
		{"Один элемент, должен быть найден", []int{1}, 1, 0},
		{"Два элемента, не должен быть найден", []int{1}, 5, -1},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.nums)

			actual := search(testCase.nums, testCase.target)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
