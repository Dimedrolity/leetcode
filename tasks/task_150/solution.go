// Package task_150
// 150. Evaluate Reverse Polish Notation
package task_150

import (
	"container/list"
	"strconv"
)

// evalRPN
// Мое решение.
// Сложность по времени O(N), по памяти O(N).
// Обратная польская нотация. Если текущий токен является числом, кладем в стек.
// Если текущий токен является оператором, то берем последние 2 числа со стека и выполняем операцию.
func evalRPN_Slice(tokens []string) int {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	stack := make([]int, 0)
	for _, t := range tokens {
		isOperator := operators[t]
		if isOperator {
			second := stack[len(stack)-1]
			first := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var val int
			switch t {
			case "+":
				val = first + second
			case "-":
				val = first - second
			case "*":
				val = first * second
			case "/":
				val = first / second
			}
			stack = append(stack, val)
		} else {
			val, err := strconv.Atoi(t)
			if err != nil {
				return -1
			}
			stack = append(stack, val)
		}
	}

	return stack[0]
}

///

// Двусвязный список из стандартной библиотеки.
func evalRPN_List(tokens []string) int {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	stack := list.New()
	for _, t := range tokens {
		isOperator := operators[t]
		if isOperator {
			second := stack.Remove(stack.Front()).(int)
			first := stack.Remove(stack.Front()).(int)

			var val int
			switch t {
			case "+":
				val = first + second
			case "-":
				val = first - second
			case "*":
				val = first * second
			case "/":
				val = first / second
			}
			stack.PushFront(val)
		} else {
			val, err := strconv.Atoi(t)
			if err != nil {
				return -1
			}
			stack.PushFront(val)
		}
	}

	return stack.Front().Value.(int)
}

///
// Стек на дженериках и двусвязном списке из стандартной библиотеки.

type Stack[v any] struct {
	l *list.List
}

func NewStack[v any]() *Stack[v] {
	return &Stack[v]{l: list.New()}
}

func (s *Stack[v]) Push(val v) {
	s.l.PushBack(val)
}

func (s *Stack[v]) Pop() v {
	last := s.l.Back()
	s.l.Remove(last)

	return last.Value.(v)
}

func evalRPN(tokens []string) int {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	stack := NewStack[int]()
	for _, t := range tokens {
		isOperator := operators[t]
		if isOperator {
			second := stack.Pop()
			first := stack.Pop()

			var val int
			switch t {
			case "+":
				val = first + second
			case "-":
				val = first - second
			case "*":
				val = first * second
			case "/":
				val = first / second
			}
			stack.Push(val)
		} else {
			val, err := strconv.Atoi(t)
			if err != nil {
				return -1
			}
			stack.Push(val)
		}
	}

	return stack.Pop()
}
