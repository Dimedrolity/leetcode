// Package task_121
// 121. Best Time to Buy and Sell Stock
package task_121

import "slices"

// maxProfit_slidingWindow
// Сложность по времени O(N), N=len(prices). По памяти O(1).
// Решение из интернета, паттерн скользящее окно.
// Правило окна: значение по левому указателю должно быть меньше значения по правому.
// Если правило выполняется, то считаем разницу и если это макс, то записываем в результат.
// Если правило нарушается, то левый указатель ставим на правый (нашли меньшую цену для покупки), правый сдвигаем на +1.
func maxProfit(prices []int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}

	var leftIndex, result int
	for rightIndex := 1; rightIndex < len(prices); rightIndex++ {
		buy := prices[leftIndex]
		sell := prices[rightIndex]

		if buy < sell {
			profit := sell - buy
			result = max(result, profit)
		} else {
			leftIndex = rightIndex
		}
	}

	return result
}

// maxProfit_My
// Собственное решение.
// Сложность по времени O(N), N=len(prices). По памяти O(1).
// На каждой итерации знаем минимальный элемент из отрезка [:i].
// Вычисляем разницу на каждой итерациии записываем максимум в результат.
// Имхо минус в том, что вычисление на каждой итерации (может быть дорогостоящая операция вместо вычитания).
// Также это не по паттерну, паттерны проще для понимания, люди их узнают.
func maxProfit_My(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	var result int
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - minPrice
		if diff > result {
			result = diff
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return result
}

// maxProfit_N2
// Первое решение. Сложность по времени O(N^2), N=len(prices).
func maxProfit_N2(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	var result int
	for i := 0; i < len(prices); i++ {
		maxToRight := slices.Max(prices[i:])
		profit := maxToRight - prices[i]

		if profit > result {
			result = profit
		}
	}

	return result
}
