// Package task_167
// 167. Two Sum II - Input Array Is Sorted
package task_167

import "slices"

// twoSum
// Мое решение.
// Сложность по времени O(N*logN), по памяти O(1).
// Идея в том, чтобы с помощью бинарного поиска найти индекс числа.
// Возможен только один индекс, так как по условию в тестах всего одно решение.
func twoSum_My(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		diff := target - numbers[i]
		j := binSearch(numbers, diff)
		if j != i && j != -1 {
			result := []int{i + 1, j + 1}
			slices.Sort(result)
			return result
		}
	}

	return nil
}

// binSearch возвращает индекс или -1.
func binSearch(nums []int, target int) int {
	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {
		middleIndex := (leftIndex + rightIndex) / 2
		if target == nums[middleIndex] {
			return middleIndex
		} else if target > nums[middleIndex] {
			leftIndex = middleIndex + 1
		} else {
			rightIndex = middleIndex - 1
		}
	}

	return -1
}

// twoSum
// Решение из интернета
// Сложность по времени O(N), по памяти O(1).
// Идея в том, чтобы использовать 2 указателя, на начало и на конце.
// Массив отсортирован.
// Если сумма по текущим указателям больше target, то сдвигаем влево правый указатель (уменьшаем сумму).
// Если сумма по текущим указателям меньше target, то сдвигаем вправо левый указатель (увеличиваем сумму).
func twoSum(numbers []int, target int) []int {
	leftIndex := 0
	rightIndex := len(numbers) - 1

	for leftIndex < rightIndex {
		sum := numbers[leftIndex] + numbers[rightIndex]
		if sum == target {
			return []int{leftIndex + 1, rightIndex + 1}
		} else if sum < target {
			leftIndex++
		} else {
			rightIndex--
		}
	}

	return nil
}
