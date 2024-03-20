// Package task_53
// 53. Maximum Subarray
package task_53

// Первая мысль - перебрать все подмассивы - это O(N^2) по времени.

// maxSubArray
// Мое решение.
// Сложность по времени O(N), по памяти O(1).
// Идея в том, чтобы добавлять элемент в сумму, но если этот элемент больше чем (сумма предыдущих + этот элемент),
// то он становится большей суммой.
//
// P.S. Это алгоритм Кадана https://en.wikipedia.org/wiki/Maximum_subarray_problem
// Можно упростить: localMaxSum = max(num, localMaxSum+num)
func maxSubArray(nums []int) int {
	localMaxSum := nums[0] // не 0, так как возможны отрицательные числа.
	globalMaxSum := localMaxSum

	for _, num := range nums[1:] {
		localMaxSum += num
		if num > localMaxSum {
			localMaxSum = num
		}
		globalMaxSum = max(globalMaxSum, localMaxSum)
	}

	return globalMaxSum
}

// maxSubArray
// Решение из интернета.
// Сложность по времени O(N), по памяти O(1).
// Идея в том, чтобы добавлять элемент в сумму, но если текущая сумма меньше нуля, то обнулять сумму.
// Имхо, решение чуть более сложное для понимания.
func maxSubArray_NotMy(nums []int) int {
	globalMaxSum := nums[0] // не 0, так как возможны отрицательные числа.
	var localMaxSum int

	for _, num := range nums {
		if localMaxSum < 0 {
			localMaxSum = 0 // если 2 отриц числа, то нет смысла их складывать, нужно взять наибольшее из них.
		}
		localMaxSum += num
		globalMaxSum = max(globalMaxSum, localMaxSum)
	}

	return globalMaxSum
}
