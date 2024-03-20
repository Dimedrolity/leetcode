// Package task_217
// 33. Search in Rotated Sorted Array
package task_33

// search
// Мое решение, без рекурсии.
// Сложность по времени O(logN), по памяти O(1).
// Аналог бинарного поиска. Условно делим массив на 2 части: от 0 до n/2 и от n/2 до n-1.
// Находим отсортированную часть, и если target находится ней (target больше начала и меньше конца),
// то в следующей итерации ищем в этой части, иначе в другой части.
// P.S. за счет проверок (nums[leftIndex] == target) и (nums[rightIndex] == target), отпадает необходимость в знаках <= >=.
func search_My(nums []int, target int) int {
	result := -1

	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {
		if nums[leftIndex] == target {
			return leftIndex
		}
		if nums[rightIndex] == target {
			return rightIndex
		}

		middleIndex := (leftIndex + rightIndex) / 2
		if nums[middleIndex] == target {
			return middleIndex
		}

		if nums[leftIndex] < nums[middleIndex] { // левая часть отсортирована
			if nums[leftIndex] < target && target < nums[middleIndex] {
				rightIndex = middleIndex - 1 // продолжение поиска target в левой части
			} else {
				leftIndex = middleIndex + 1
			}
		} else { // правая часть отсортирована
			if nums[middleIndex] < target && target < nums[rightIndex] {
				leftIndex = middleIndex + 1 // продолжение поиска target в правой части
			} else {
				rightIndex = middleIndex - 1
			}
		}
	}

	return result
}

// search
// Решение из интернета, доработанное.
// Сложность по времени O(logN), по памяти O(1).
// Аналог бинарного поиска. Условно делим массив на 2 части: от 0 до n/2 и от n/2 до n-1.
// Находим отсортированную часть, и если target находится ней (target больше начала и меньше конца),
// то в следующей итерации ищем в этой части, иначе в другой части.
// Идем пока middle не будет указывать на target.
func search(nums []int, target int) int {
	result := -1

	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {
		middleIndex := (leftIndex + rightIndex) / 2
		if nums[middleIndex] == target {
			return middleIndex
		}

		if nums[leftIndex] <= nums[middleIndex] { // левая часть отсортирована
			if nums[leftIndex] <= target && target <= nums[middleIndex] {
				rightIndex = middleIndex - 1 // продолжение поиска target в левой части
			} else {
				leftIndex = middleIndex + 1
			}
		} else { // правая часть отсортирована
			if nums[middleIndex] <= target && target <= nums[rightIndex] {
				leftIndex = middleIndex + 1 // продолжение поиска target в правой части
			} else {
				rightIndex = middleIndex - 1
			}
		}
	}

	return result
}
