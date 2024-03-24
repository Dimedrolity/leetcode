// Package task_11
// 11. Container With Most Water
package task_11

// maxArea
// Мое решение
// Сложность по времени O(n^2), по памяти O(1).
// Перебор всех чисел. Не эффективное по времени решение.
func maxArea_My(height []int) int {
	var current, result int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			w := j - i
			h := min(height[i], height[j])
			current = w * h
			result = max(result, current)
		}
	}

	return result
}

// maxArea
// Решение из интернета.
// Сложность по времени O(n), по памяти O(1).
// Идея в том, чтобы поставить 2 указателя, на начало и на конец.
// Сдвигать тот указатель, число по которому меньше.
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var result int

	for left < right {
		w := right - left
		h := min(height[left], height[right])
		area := h * w
		result = max(result, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}
