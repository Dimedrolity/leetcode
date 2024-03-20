// Package task_242
// 242. Valid Anagram
package task_242

// containsDuplicate
// Мое решение.
// Сложность по времени O(s+t), по памяти O(s+t).
// Идея в том, чтобы посчитать частоту встречаемости символов каждой строки и сравнить частотные словари.
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := wordList(s)
	tMap := wordList(t)

	return mapsEqual(sMap, tMap)
}

func wordList(s string) map[rune]int {
	wl := make(map[rune]int)
	for _, letter := range s {
		wl[letter] += 1
	}
	return wl
}

func mapsEqual(m1, m2 map[rune]int) bool {
	for letter, count := range m1 {
		if m2[letter] != count {
			return false
		}
	}

	return true
}
