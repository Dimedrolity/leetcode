package task_153

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	nums     []int
	expected int
}

/*
Example 1:

Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.
Example 2:

Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
Example 3:

Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.
*/

func TestFindMin(t *testing.T) {
	testCases := []TestCase{
		{"Сдвиг на 1", []int{7, 1, 2, 3, 4, 5, 6}, 1},
		{"Сдвиг на 2", []int{6, 7, 1, 2, 3, 4, 5}, 1},
		{"Сдвиг на 3", []int{5, 6, 7, 1, 2, 3, 4}, 1},
		{"Сдвиг на 4", []int{4, 5, 6, 7, 1, 2, 3}, 1},
		{"Сдвиг на 5", []int{3, 4, 5, 6, 7, 1, 2}, 1},
		{"Сдвиг на 6", []int{2, 3, 4, 5, 6, 7, 1}, 1},
		{"Сдвиг на 7", []int{1, 2, 3, 4, 5, 6, 7}, 1},
		{"Два элемента, сдвиг на 1", []int{2, 1}, 1},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.nums)

			actual := findMin(testCase.nums)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
