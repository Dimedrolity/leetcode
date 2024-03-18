package task_gofunc

// GoFunc Mock Interview
// https://gofunc.ru/talks/319d7fc0852d464cb899f2b4ee09848b/?referer=/schedule/days/

// isStrValid проверяет, что последовательность скобок валидна.
//
// При встрече закрывающей скобки (symbol in closedBrackets.keys()) нужно делать stack.Peak() и если на вершине стека есть открывающая скобка этого же типа (Peak() == closedToOpen[symbol]),
// то делаем pop, так как скобки совпали, идем дальше.
//
// Пробный проход по строке ({asdf})
// stack = (
// stack = (, {
// stack = (, {, }
// stack = (
// stack = (, )
// stack = пусто
func isStrValid(s string) bool {
	openBrackets := map[rune]bool{'(': true, '{': true, '[': true}
	closedBrackets := map[rune]bool{')': true, '}': true, ']': true}
	closedToOpen := map[rune]rune{')': '(', '}': '{', ']': '['}
	var stack []rune

	for _, symbol := range s {
		if _, isOpenBracket := openBrackets[symbol]; isOpenBracket {
			stack = append(stack, symbol) // Push
		} else if _, isCloseBracket := closedBrackets[symbol]; isCloseBracket {
			if len(stack) == 0 {
				return false
			}
			n := len(stack) - 1
			top := stack[n]
			if top == closedToOpen[symbol] {
				stack = stack[:n] // Pop
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
