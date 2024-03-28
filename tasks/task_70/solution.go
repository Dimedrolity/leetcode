// Package task_70
// 70. Climbing Stairs
package task_70

// climbStairs
// Мое решение.
// Сложность по времени O(N), по памяти O(N).
// Идея в том, чтобы хранить информацию о предыдущих ступеньках.
// Для каждой ступеньки хранится кол-во вариантов, как подняться на нее.
// На каждую ступеньку начиная с третьей, можно подняться одним шагом с последней или двумя шагами с предпоследней.
func climbStairs_My(n int) int {
	stair := make(map[int]int, n+1)
	stair[1] = 1
	stair[2] = 2
	for i := 3; i < n+1; i++ {
		stair[i] += stair[i-2]
		stair[i] += stair[i-1]
	}

	return stair[n]
}

// climbStairs
// Решение из интернета
// Сложность по времени O(N), по памяти O(1).
// Идея в том, чтобы найти (n-1)-ое число Фибоначчи.
// Это будет эквивалентно тому, что до посчитать варианты от n до 0 ступеньки.
func climbStairs(n int) int {
	one, two := 1, 1
	for i := 0; i < n; i++ {
		tmpOne := one
		one = two
		two = tmpOne + two
	}

	return one
}
