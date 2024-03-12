// Package task_217
// 217. Contains Duplicate
package task_217

// containsDuplicate
// Мое решение.
// Сложность по времени O(N), по памяти O(N).
// Set содержит униккальные значения без дубликатов.
// Если в массиве есть дубликаты, то кол-во элементов в сете будет меньше на кол-во повторяющихся элементов.
func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	return len(set) < len(nums)
}
