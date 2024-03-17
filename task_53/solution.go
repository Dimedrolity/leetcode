// Package task_53
// 53. Maximum Subarray
package task_53

// Первая мысль - перебрать все подмассивы - это O(N^2) по времени.

// maxSubArray
// Мое решение.
// Сложность по времени O(N), по памяти O(1).
// Идея в том, чтобы добавлять элемент в сумму, но если этот элемент больше чем (сумма предыдущих + этот элемент),
// то он становится большей суммой.
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	localMaxSum := nums[0] // не 0, так как возможны отрицательные числа.
	globalMaxSum := localMaxSum
	for i := 1; i < len(nums); i++ {
		localMaxSum += nums[i]
		if nums[i] > localMaxSum {
			localMaxSum = nums[i]
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
	if len(nums) == 1 {
		return nums[0]
	}

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
