package task_217_contains_duplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	nums     []int
	expected bool
}

func TestContainsDuplicate(t *testing.T) {
	testCases := []TestCase{
		{"Дублирующийся элемент присутствует единожды", []int{1, 2, 3, 1}, true},
		{"Дублирующийся элемент отсутствует", []int{1, 2, 3, 4}, false},
		{"Дублирующийся элемент присутствует несколько раз", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{"Массив из одного элемента", []int{1}, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := containsDuplicate(testCase.nums)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
