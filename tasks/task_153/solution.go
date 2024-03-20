// Package task_153
// 153. Find Minimum in Rotated Sorted Array
package task_153

import "math"

// findMin
// Мое решение.
// Сложность по времени O(logN), по памяти O(1).
// Идея в том, чтобы условно разделить массив на 2 части от 0 до n/2 и от n/2 до n-1.
// По этим трем числам определить, где находятся 2 элемента в убывающей подпоследовательности (первый больше второго).
// Аналог бинарного поиска.
func findMin_My(nums []int) int {
	if len(nums) == 0 {
		return math.MaxInt32
	}
	if len(nums) == 1 {
		return nums[0]
	}

	leftIndex := 0
	middleIndex := len(nums) / 2
	rightIndex := len(nums) - 1

	if nums[leftIndex] < nums[middleIndex] {
		if nums[middleIndex] < nums[rightIndex] {
			// l < m < r
			// значит мин в начале
			return nums[leftIndex]
		}

		// l < m > r
		// мин между m и r
		return min(nums[leftIndex], nums[rightIndex], findMin(nums[middleIndex:rightIndex]))
	}
	// l > m < r
	// мин между l и m
	// По условию не может быть l > m > r - это убывающая последовательность.
	return min(nums[middleIndex], findMin(nums[leftIndex:middleIndex]))
}

// findMin
// Решение из интернета, без рекурсии, 2 указателя.
// Сложность по времени O(logN), по памяти O(1).
// Идея в том, чтобы условно разделить массив на 2 части от 0 до n/2 и от n/2 до n-1.
// По этим трем числам определить, где находятся 2 элемента в убывающей подпоследовательности (первый больше второго).
// Аналог бинарного поиска.
func findMin(nums []int) int {
	result := nums[0]

	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {
		if nums[leftIndex] < nums[rightIndex] {
			result = min(result, nums[leftIndex])
			break
		}

		middleIndex := (leftIndex + rightIndex) / 2
		result = min(result, nums[middleIndex])

		if nums[leftIndex] <= nums[middleIndex] {
			leftIndex = middleIndex + 1
		} else {
			rightIndex = middleIndex - 1
		}
	}

	return result
}
