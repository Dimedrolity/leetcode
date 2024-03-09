package task_121

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestMaxProfit_Descending возрастающая последовательность.
func TestMaxProfit_Ascending(t *testing.T) {
	actual := maxProfit([]int{1, 2, 3, 4})
	expected := 3
	assert.Equal(t, expected, actual)
}

// TestMaxProfit_MaxToLeftOfMin максимальный левее минимального, поэтому не учитывается.
func TestMaxProfit_MaxToLeftOfMin(t *testing.T) {
	actual := maxProfit([]int{7, 1, 5, 3, 6, 4})
	expected := 5
	assert.Equal(t, expected, actual)
}

// TestMaxProfit_Descending убывающая последовательность.
func TestMaxProfit_Descending(t *testing.T) {
	actual := maxProfit([]int{7, 6, 4, 3, 1})
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestMaxProfit_Nil(t *testing.T) {
	actual := maxProfit(nil)
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestMaxProfit_Empty(t *testing.T) {
	actual := maxProfit(nil)
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestMaxProfit_Single(t *testing.T) {
	actual := maxProfit([]int{1})
	expected := 0
	assert.Equal(t, expected, actual)
}

// TestMaxProfit_MinInTheEnd
// Минимальный элемент в конце массива, поэтому не учитывается.
func TestMaxProfit_MinInTheEnd(t *testing.T) {
	actual := maxProfit([]int{2, 4, 1})
	expected := 2
	assert.Equal(t, expected, actual)
}
