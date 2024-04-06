// Package task_217_contains_duplicate
// 217. Contains Duplicate
package task_217_contains_duplicate

import (
	"strconv"
)

// isPalindrome
// Мое решение без преобразования числа в строку.
// Идея в том, чтобы вычислить разряды числа и сравнить разряды с помощью двух указателей.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	r := radix(x)

	left, right := 0, len(r)-1
	for left < right {
		if r[left] != r[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func radix(x int) []int {
	r := make([]int, 0)

	for x > 0 {
		r = append(r, x%10)
		x = x / 10
	}

	return r
}

// isPalindrome
// Мое решение с преобразованием числа в строку.
// Идея в том, чтобы сравнить байты с помощью двух указателей.
func isPalindrome_Strconv(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
