// Package task_128
// 128. Longest Consecutive Sequence
package task_128

import "slices"

// longestConsecutive_My_N
// Мое решение.
// Сложность по времени O(N), по памяти O(N).
// Идея в том, чтобы записать числа в Set, пройтись по числам,
// на каждой итерации проверка, если есть элемент меньше на 1, то добавляем в последовательность текущий элемент;
// аналогично, если есть элемент больше на 1, то последовательностью будут текущие элементы последовательность элементов больше на 1.
// Довольно сложное, но рабочее решение.
func longestConsecutive_My_N(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	numToSeq := make(map[int][]int)
	var longestSequence []int
	for num, _ := range set {
		if seq, ok := numToSeq[num-1]; ok {
			numToSeq[num] = append(seq, num)

			// обновление начала и конца последовательности.
			first := numToSeq[num][0]
			numToSeq[first] = numToSeq[num]
			last := numToSeq[num][len(numToSeq[num])-1]
			numToSeq[last] = numToSeq[num]
		} else {
			numToSeq[num] = []int{num}
		}
		if seq, ok := numToSeq[num+1]; ok {
			numToSeq[num] = append(numToSeq[num], seq...)

			// обновление начала и конца последовательности.
			first := numToSeq[num][0]
			numToSeq[first] = numToSeq[num]
			last := numToSeq[num][len(numToSeq[num])-1]
			numToSeq[last] = numToSeq[num]
		}
		if len(numToSeq[num]) > len(longestSequence) {
			longestSequence = numToSeq[num]
		}
	}

	return len(longestSequence)
}

// O(N*logN) с помощью set, sort и подсчета кол-ва элементов.
func longestConsecutive_My_NlogN(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	uniqueNums := make([]int, 0, len(set))
	for num, _ := range set {
		uniqueNums = append(uniqueNums, num)
	}

	slices.Sort(uniqueNums)
	var result int

	currentLongest := 1
	result = currentLongest
	for i := 0; i < len(uniqueNums)-1; i++ {
		if uniqueNums[i]+1 == uniqueNums[i+1] {
			currentLongest += 1
		} else {
			currentLongest = 1
		}
		result = max(result, currentLongest)
	}

	return result
}

// longestConsecutive
// Решение из интернета.
// Сложность по времени O(N), по памяти O(N).
// Идея в том, чтобы записать числа в Set, пройтись по числам, и на каждой итерации проверять, является ли число началом последовательности
// (число является началом последовательности, если нет числа меньше на 1).
// Если оно является началом последовательности, то собираем всю последовательность с помощью цикла, проверяя, есть ли элемент на 1 больше.
func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	result := 1

	for num := range set {
		if set[num-1] {
			continue
		}

		length := 1
		for set[num+length] {
			length++
		}

		result = max(result, length)
	}

	return result
}
