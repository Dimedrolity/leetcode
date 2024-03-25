// Package task_125
// 125. Valid Palindrome
package task_125

import (
	"strings"
	"unicode"
)

// isPalindrome
// Мое решение.
// Сложность по времени O(N), по памяти O(N).
// Идея в том, чтобы отфильтровать только цифры и буквы, привести к нижнему регистру и сравнить с реверснутой строкой.
func isPalindrome_MyFirst(s string) bool {
	b := strings.Builder{}
	for _, r := range s {
		if ('0' <= r && r <= '9') || ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
			b.WriteRune(r)
		}
	}
	alphanumeric := b.String()
	alphanumeric = strings.ToLower(alphanumeric)
	rev := reverse(alphanumeric)
	return alphanumeric == rev
}

func reverse(s string) string {
	b := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}

	return b.String()
}

func isAlphanumeric(s string) string {
	b := strings.Builder{}
	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			b.WriteRune(r)
		}
	}
	alphanum := b.String()

	return alphanum
}

// isPalindrome
// Мое решение после рефакторинга.
// Сложность по времени O(N), по памяти O(N).
// Идея в том, чтобы отфильтровать только цифры и буквы, привести к нижнему регистру и сравнить с реверснутой строкой.
func isPalindrome(s string) bool {
	alphanum := isAlphanumeric(s)
	alphanum = strings.ToLower(alphanum)
	rev := reverse(alphanum)

	return alphanum == rev
}
