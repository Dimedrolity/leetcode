// Package task_122
// 122. Best Time to Buy and Sell Stock II
package task_122

// maxProfit
// Решение из интернета.
// Сложность по времени O(N), N=len(prices). По памяти O(1).
// Более простое решение.
// Покупаем и продаем каждый раз, когда цена увеличилась
// (на каждой итерации добавляем разницу между числами, если есть рост цены).
func maxProfit(prices []int) int {
	var profit int

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

// maxProfit
// Собственное решение.
// Сложность по времени O(N), N=len(prices). По памяти O(1).
// Идея в том, чтобы продавать акцию перед тем, как цена понижается.
// Скользящее окно. В начале каждой итерации left < right.
// Левый указатель указывает на цену покупки, правый - на цену продажи.
// Если на текущей итерации понимаем, что акция падает (профит уменьшается),
// то продаем (суммируем к результату), двигаем указатели и снова ищем профит.
func maxProfit_My(prices []int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}

	var leftIndex, localMaxProfit, result int
	rightIndex := 1
	for rightIndex < len(prices) {
		buy := prices[leftIndex]
		sell := prices[rightIndex]
		profit := sell - buy

		if profit > localMaxProfit {
			localMaxProfit = profit
		} else {
			result += localMaxProfit
			localMaxProfit = 0
			leftIndex = rightIndex
		}

		rightIndex++
	}

	// Суммируем на случай возрастающей последовательности.
	if localMaxProfit != 0 {
		result += localMaxProfit
	}

	return result
}
