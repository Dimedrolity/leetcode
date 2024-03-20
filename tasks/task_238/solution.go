// Package task_238
// 238. Product of Array Except Self
package task_238

// productExceptSelf
// Мое решение с делением.
// Сложность по времени O(N), по памяти O(1) если не учитывать массив результата.
// Идея в том, чтобы получить произведение всех элементов, пройтись по каждому элементу и разделить произведение на него.
// Если есть один ноль в массиве чисел, то только на его месте будет product, на остальных местах нули.
// Если есть 2 ноля и более, то весь результирующей массив будет состоять из нулей.
func productExceptSelf_division(nums []int) []int {
	product := 1
	zeroCount := 0
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			product *= num
		}
	}

	result := make([]int, len(nums))

	if zeroCount == 0 {
		for i, num := range nums {
			if num != 0 {
				result[i] = product / num
			}
		}
	} else if zeroCount == 1 {
		for i, num := range nums {
			if num == 0 {
				result[i] = product
			}
		}
	}

	return result
}

// productExceptSelf
// Мое решение без деления.
// Сложность по времени O(N^2), по памяти O(1) если не учитывать массив результата.
// Идея в том, чтобы пройтись по массиву и для каждого элемента умножать все элементы результирующего на него, кроме nums[i].
func productExceptSelf_N2(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		result[i] = 1
	}

	for i, num := range nums {
		for j := 0; j < len(result); j++ {
			if j != i {
				result[j] = result[j] * num
			}
		}
	}

	return result
}

// productExceptSelf
// Идея "Product prefix" из интернета, мое решение.
// Сложность по времени O(N), по памяти O(N) если не учитывать массив результата.
// Идея в том, чтобы для каждого элемента получить результат как перемножение элементов до него на перемножение после него.
// Для реализации потребуется вычислить перемножения всех префиксов и перемножения всех постфиксов.
// Результатом будет перемножение элемента из перемножения префиксов на элемент из перемножения постфиксов.
func productExceptSelf(nums []int) []int {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = 1 // До 0 индекса нет префикса, поэтому 1 для безопасного перемножения.
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] * nums[i-1]
	}

	postfixSum := make([]int, len(nums))
	postfixSum[len(nums)-1] = 1 // После последнего индекса нет постфикса, поэтому 1 для безопасного перемножения.
	for i := len(nums) - 2; i >= 0; i-- {
		postfixSum[i] = postfixSum[i+1] * nums[i+1]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		result[i] = prefixSum[i] * postfixSum[i]
	}

	return result
}
