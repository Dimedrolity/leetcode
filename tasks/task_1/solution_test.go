package task_121

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum_InStart(t *testing.T) {
	actual := twoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}
	assert.Equal(t, expected, actual)
}

func TestTwoSum_InEnd(t *testing.T) {
	actual := twoSum([]int{3, 2, 4}, 6)
	expected := []int{1, 2}
	assert.Equal(t, expected, actual)
}

func TestTwoSum_TwoElements(t *testing.T) {
	actual := twoSum([]int{3, 3}, 6)
	expected := []int{0, 1}
	assert.Equal(t, expected, actual)
}
