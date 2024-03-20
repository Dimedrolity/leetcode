// Package task_152
// 152. Maximum Product Subarray
package task_152

import "slices"

// Нелегкая задача.

// maxProduct
// Мое решение.
// Сложность по времени O(N^2), по памяти O(1).
func maxProduct_My(nums []int) int {
	globalMaxProduct := nums[0]

	for i := 0; i < len(nums); i++ {
		localMaxProduct := nums[i]
		globalMaxProduct = max(globalMaxProduct, localMaxProduct)

		for j := i + 1; j < len(nums); j++ {
			localMaxProduct *= nums[j]
			globalMaxProduct = max(globalMaxProduct, localMaxProduct)
		}
	}
	return globalMaxProduct
}

// maxProduct
// Решение из интернета.
// Сложность по времени O(N), по памяти O(1).
// Идея в том, чтобы вычислить макс и мин подмассив из первых двух элементов.
// Пройтись по массиву и на каждой итерации вычислять максимум и минимум от текущего элемента и текущих мин и макс.
func maxProduct(nums []int) int {
	globalMaxProduct := slices.Max(nums)
	localMinProduct, localMaxProduct := 1, 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			localMinProduct, localMaxProduct = 1, 1 // нейтральные значения для умножения
		}

		tmp := localMinProduct
		localMinProduct = min(nums[i], nums[i]*localMinProduct, nums[i]*localMaxProduct)
		localMaxProduct = max(nums[i], nums[i]*tmp, nums[i]*localMaxProduct)
		// не подойдет min(localMaxProduct, nums[i]*tmp, nums[i]*localMaxProduct), потому
		// допустим массив [2, 3, -1, 4] и на i=1 localMinProduct должен быть 3, а не 2, так как при дальнейшем вычислении
		// значение 2 уже не учитывается, иначе это будет перепрыжка через 1 элемент (тройку).

		globalMaxProduct = max(globalMaxProduct, localMaxProduct)
	}
	return globalMaxProduct
}

// maxProduct
// Не завершенное решение.
// Сложность по времени O(N), по памяти O(N).
// Идеи и замечания
// Числа всего лишь от -10 до 10. Можно использовать число как ключ словаря. Это для того, чтобы вычисление всех префиксов и постфиксов входило в int.
// При умножении двух отрицательных чисел, получается большое положительное.
// Что если построить массивы перемножений префиксов и постфиксов?
// Если только положительные числа, то макс произведение будет из всех чисел массива.
// Если одно отриц число, то находим произв слева от него и справа от него, берем максимум.
// Если кол-во отриц чисел четно, то перемножаем все числа.
// Если 3 отриц числа? Перемножение двух отриц дает положительное, поэтому нужно выбрать 2 из трех отриц чисел и макс произв из положительных.
// Если есть одно число 0, то вычисляем слева и справа. Берем max(0, left, right)
// Если нет нулей и кол-во отрицательных четно, то перемножаем все.
func maxProduct_NotDone(nums []int) int {
	var negativeCount, zeroCount int
	for _, num := range nums {
		if num < 0 {
			negativeCount++
		}
		if num == 0 {
			zeroCount++
		}
	}

	if zeroCount == 0 && negativeCount%2 == 0 {
		product := 1
		for _, num := range nums {
			product *= num
		}
		return product
	}

	product := slices.Max(nums)

	bestPrefixes := make([]int, len(nums))
	bestPrefixes[0] = nums[0]
	var leftIndex int
	for i := 1; i < len(bestPrefixes); i++ {
		bestPrefixes[i] = bestPrefixes[i-1] * nums[i]
		// while и проверка на left < len(nums) ??
		if nums[leftIndex] == 0 {
			leftIndex++
		} else if bestPrefixes[i]/nums[leftIndex] > bestPrefixes[i] {
			bestPrefixes[i] = bestPrefixes[i] / nums[leftIndex]
			leftIndex++
		}
	}
	product = max(product, slices.Max(bestPrefixes))

	bestPostfixes := make([]int, len(nums))
	bestPostfixes[len(bestPostfixes)-1] = nums[len(nums)-1]
	rightIndex := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		bestPostfixes[i] = bestPostfixes[i+1] * nums[i]
		if nums[rightIndex] == 0 {
			rightIndex--
		} else if bestPrefixes[i]/nums[rightIndex] > bestPrefixes[i] {
			bestPrefixes[i] = bestPrefixes[i] / nums[rightIndex]
			rightIndex--
		}
	}
	product = max(product, slices.Max(bestPostfixes))

	return product
}
