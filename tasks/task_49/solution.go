// Package task_49
// 49. Group Anagrams
package task_49

import (
	"sort"
)

// groupAnagrams
// Мое решение.
// Сложность по времени O(N*logN), по памяти O(N), где N - суммарная длина строк массива strs.
// Если длина всех строк равна M и количество строк равно P, то сложность по времени O(M*P*log(M*P)).
// Идея в том, чтобы сгруппировать строки по отсортированной строке и вернуть группы.
func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		if anagrams[sortedStr] == nil {
			anagrams[sortedStr] = []string{}
		}
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	groups := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		groups = append(groups, group)
	}

	return groups
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i int, j int) bool { return runes[i] < runes[j] })
	sortedStr := string(runes)
	return sortedStr
}

// Есть решение за O(M*P), оно чуть сложнее для понимания и кода будет больше. Не стал кодить.
