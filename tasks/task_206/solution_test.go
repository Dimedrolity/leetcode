package task_206

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	head     *ListNode
	expected *ListNode
}

func sliceToList(nums []int) *ListNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

func TestReverseList(t *testing.T) {

	testCases := []TestCase{
		{"Несколько чисел", sliceToList([]int{1, 2, 3, 4, 5}), sliceToList([]int{5, 4, 3, 2, 1})},
		{"Два числа", sliceToList([]int{1, 2}), sliceToList([]int{2, 1})},
		{"Одно число", sliceToList([]int{1}), sliceToList([]int{1})},
		{"Пустой список", sliceToList([]int{}), sliceToList(([]int{}))},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := reverseList(testCase.head)
			curActual := actual
			curExpected := testCase.expected

			for curActual != nil && curExpected != nil {
				assert.Equal(t, curExpected.Val, curActual.Val)
				curExpected = curExpected.Next
				curActual = curActual.Next
			}
			assert.Nil(t, curActual)
			assert.Nil(t, curExpected)
		})
	}
}
