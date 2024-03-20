package task_53

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

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
*/

func TestMaxSubArray(t *testing.T) {
	testCases := []TestCase{
		{"Макс подмассив в середине", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"Массив из 1-го элемента", []int{1}, 1},
		{"Макс подмассив состоит из всех элементов", []int{5, 4, -1, 7, 8}, 23},
		{"Массив из отрицательных чисел", []int{-5, -4, -1, -7, -8}, -1},
		{"Возрастающая последовательность", []int{1, 2, 3, 4}, 10},
		{"Убывающая последовательность", []int{4, 3, 2, 1}, 10},
		{"Убывающая последовательность", []int{-3, -1, -2}, -1},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.nums)

			actual := maxSubArray(testCase.nums)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
