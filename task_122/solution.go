// Package task_122
// 122. Best Time to Buy and Sell Stock II
package task_122

func maxProfit(prices []int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}

	var leftIndex, localMaxProfit, result int
	rightIndex := 1
	for rightIndex < len(prices) {
		buy := prices[leftIndex]
		sell := prices[rightIndex]
		profit := sell - buy

		if profit > 0 {
			if profit < localMaxProfit {
				result += localMaxProfit
				localMaxProfit = 0
				leftIndex = rightIndex
			} else {
				localMaxProfit = profit
			}
		} else {
			result += localMaxProfit
			localMaxProfit = 0
			leftIndex = rightIndex
		}
		rightIndex++
	}

	if localMaxProfit != 0 {
		result += localMaxProfit
	}

	return result
}
