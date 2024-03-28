package task_70

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	n        int
	expected int
}

func TestClimbStairs(t *testing.T) {
	testCases := []TestCase{
		{"Вторая ступенька", 2, 2},
		{"Третья ступенька", 3, 3},
		{"Пятая ступенька", 5, 8},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := climbStairs(testCase.n)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
