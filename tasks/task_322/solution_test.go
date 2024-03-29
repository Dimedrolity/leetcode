package task_322

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	coins    []int
	amount   int
	expected int
}

func TestCoinChange(t *testing.T) {
	testCases := []TestCase{
		{"Дважды одна монета и единожды другая монета", []int{1, 2, 5}, 11, 3},
		{"Не складывается в нужную сумму", []int{2}, 3, -1},
		{"Для суммы 0 нужно 0 монет", []int{1}, 0, 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := coinChange(testCase.coins, testCase.amount)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
