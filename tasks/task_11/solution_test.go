package task_11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	height   []int
	expected int
}

func TestMaxArea(t *testing.T) {
	testCases := []TestCase{
		{"Большая площадь не по макс-й высоте", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"Две высоты", []int{1, 1}, 1},
		{"Возрастающая последовательность", []int{1, 2, 3, 4, 5, 6}, 9},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := maxArea(testCase.height)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
