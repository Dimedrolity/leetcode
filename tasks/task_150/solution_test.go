package task_150

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	tokens   []string
	expected int
}

func TestEvalRPN(t *testing.T) {
	testCases := []TestCase{
		{"Выражение с приоритетом на сложение (2+1)*3", []string{"2", "1", "+", "3", "*"}, 9},
		{"Выражение с приоритетом на деление 4+(13/5)", []string{"4", "13", "5", "/", "+"}, 6},

		/*
			((10 * (6 / ((9 + 3) * -11))) + 17) + 5
			= ((10 * (6 / (12 * -11))) + 17) + 5
			= ((10 * (6 / -132)) + 17) + 5
			= ((10 * 0) + 17) + 5
			= (0 + 17) + 5
			= 17 + 5
			= 22
		*/
		{"Выражение с делением и умножением", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := evalRPN(testCase.tokens)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
