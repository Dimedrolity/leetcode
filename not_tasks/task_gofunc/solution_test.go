package task_gofunc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	str      string
	expected bool
}

func TestIsStrValid(t *testing.T) {
	testCases := []TestCase{
		{"Один вид скобок", "(qwer)", true},
		{"Три вида скобок", "({[abc]})", true},
		{"Нет закрывающей скобки", "({)", false},
		{"Нет открывающей скобки", ")", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := isStrValid(testCase.str)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
