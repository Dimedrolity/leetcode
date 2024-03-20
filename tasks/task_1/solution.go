// Package task_121
// 121. Best Time to Buy and Sell Stock
package task_121

import "slices"

// twoSum
// Мое решение.
// Сложность по времени O(N^2), по памяти O(1).
func twoSum_My_N2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// twoSum
// После подсказки №2 про поиск числа равного `target - nums[i]`.
// Сложность по времени O(N*logN), по памяти O(N).
func twoSum_BinSearch(nums []int, target int) []int {
	numToIndexes := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		if numToIndexes[nums[i]] == nil {
			numToIndexes[nums[i]] = make([]int, 0)
		}
		numToIndexes[nums[i]] = append(numToIndexes[nums[i]], i)
	}

	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {
		localTarget := target - nums[i]
		j, found := slices.BinarySearch(nums, localTarget)
		if !found {
			continue
		}
		if nums[i] == nums[j] {
			return numToIndexes[nums[i]]
		} else {
			sourceI := numToIndexes[nums[i]][0]
			sourceJ := numToIndexes[nums[j]][0]
			return []int{sourceI, sourceJ}
		}
	}

	return nil
}

// twoSum
// После подсказки №3 про hashmap стало понятно, что сортировка и бин поиск тут ни к чему.
// Сложность по времени O(N), по памяти O(N).
func twoSum_Map(nums []int, target int) []int {
	numToIndexes := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		if numToIndexes[nums[i]] == nil {
			numToIndexes[nums[i]] = make([]int, 0)
		}
		numToIndexes[nums[i]] = append(numToIndexes[nums[i]], i)
	}

	for i := 0; i < len(nums); i++ {
		localTarget := target - nums[i]
		indexes, found := numToIndexes[localTarget]
		if !found {
			continue
		}
		// Два одинаковых числа дают в сумме target
		if len(indexes) == 2 {
			return indexes
		} else {
			// Само с собой не нужно складывать
			if indexes[0] == i {
				continue
			}
			return []int{i, indexes[0]}
		}
	}

	return nil
}

// twoSum
// Решение из интеренета.
// Сложность по времени O(N), по памяти O(N).
// Идея в том, чтобы в одном цикле и заполнять словарь и искать элемент.
// Таким образом и не нужно будет хранить массив индексов в качестве значения как в моем решении.
func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]

		if j, found := numToIndex[diff]; found {
			return []int{j, i}
		}

		numToIndex[nums[i]] = i
	}

	return nil
}
