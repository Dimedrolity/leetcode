// Package task_15
// 15. 3Sum
package task_15

import "slices"

// threeSum
// Решение из интернета.
// Сложность по времени O(N^2), по памяти O(logN) (pdqsort).
// Идея в том, чтобы отсортировать массив и использовать 2 указателя, на начало и на конце, как при решении TwoSum II.
// Проходимся по массиву, для каждого элемента делаем TwoSum II для поиска суммы равной нулю.
// Если сумма трех элементов больше нуля, то сдвигаем влево правый указатель (уменьшаем сумму).
// Если сумма трех элементов меньше нуля, то сдвигаем вправо левый указатель (увеличиваем сумму).
// Если элемент nums[i] повторяется, то пропускаем, аналогично пропускаем если повторяется элемент nums[leftIndex].
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		leftIndex, rightIndex := i+1, len(nums)-1
		for leftIndex < rightIndex {
			sum := nums[i] + nums[leftIndex] + nums[rightIndex]
			if sum > 0 {
				rightIndex--
			} else if sum < 0 {
				leftIndex++
			} else {
				result = append(result, []int{nums[i], nums[leftIndex], nums[rightIndex]})
				leftIndex++
				for leftIndex < rightIndex && nums[leftIndex] == nums[leftIndex-1] {
					leftIndex++
				}
			}
		}
	}

	return result
}
