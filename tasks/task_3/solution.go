// Package task_3
// 3. Longest Substring Without Repeating Characters
package task_3

// lengthOfLongestSubstring
// Мое решение.
// Сложность по времени O(N), по памяти O(N).
// Идея в том, чтобы реализовать паттерн sliding window.
// Создать 2 указателя, левый на начало уникальной последовательности, правый сразу после левого (является переменной цикла).
// Хранение текущей уникальной последовательности в map, ключи - элементы, значения - индексы.
func lengthOfLongestSubstring_My(s string) int {
	runes := []rune(s)
	if len(runes) == 0 {
		return 0
	}
	if len(runes) == 1 {
		return 1
	}
	left := 0
	m := map[rune]int{runes[left]: left} // rune to index
	result := 0
	for right := 1; right < len(runes); right++ {
		if _, ok := m[runes[right]]; !ok { // если нет в словаре, то записываем
			m[runes[right]] = right
		} else { // если есть в словаре, то сдвигаем левый указатель окна до тех пор, пока не встретится текущий символ.
			// и потом еще +1, чтобы перейти на новый символ, так как символ runes[right] уже есть в уникальной последовательности.

			// Не обязательная ветка, граничный случай. Навел на мысль сделать общий случай (ниже).
			//if runes[right] == runes[right-1] {
			//	left = right
			//	m = map[rune]int{runes[left]: left}
			//} else {
			for runes[left] != runes[right] {
				delete(m, runes[left])
				left++
			}
			left++
			m[runes[right]] = right
			//}
		}

		result = max(result, len(m))
	}

	return result
}

// lengthOfLongestSubstring
// Решение из интернета
// Сложность по времени O(N), по памяти O(N).
// Идея в том, чтобы реализовать паттерн sliding window.
// Создать 2 указателя, левый на начало уникальной последовательности, правый сразу после левого (является переменной цикла).
// Правила окна: содержит уникальную последовательность.
// Если текущий символ уже есть в окне (дублируется), то смещаем левый указатель пока не переступим через такой же символ как текущий.
// Хранение текущей уникальной последовательности в set, чтобы проверять, содержит ли окно текущий символ.
func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)

	left := 0
	set := map[rune]bool{}
	result := 0
	for right := 0; right < len(runes); right++ {
		// проверка по правому, а удаление по левому, гениально!
		for set[runes[right]] {
			delete(set, runes[left])
			left++
		}
		set[runes[right]] = true
		result = max(result, len(set))
	}

	return result
}
