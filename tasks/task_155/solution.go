// Package task_155
// 155. Min Stack
package task_155

/*
Решил с подсказкой:
Consider each node in the stack having a minimum value. (Credits to @aakarshmadhavan)

Реализация стека на основе двусвязного списка.
Для определения минимального элемента, в каждом узле хранится минимальное значение всего стека на текущий момент.
При добавлении первого элемента минимальным значением будет само значение.
При добавлении последующих элементов минимальным значением будет минимум от добавляемого числа и минимального числа в хвосте.
*/

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
	Min   int // мин элемент на текущий момент
}

// MinStack это стек на основе двусвязного списка.
type MinStack struct {
	tail *Node // хвост списка
}

func Constructor() MinStack {
	return MinStack{}

}

func (ms *MinStack) Push(val int) {
	if ms.tail == nil {
		ms.tail = &Node{Value: val, Min: val}
		return
	}
	minVal := min(val, ms.tail.Min)
	node := &Node{Value: val, Min: minVal}
	node.Prev = ms.tail
	ms.tail.Next = node
	ms.tail = node
}

func (ms *MinStack) Pop() {
	if ms.tail.Prev == nil {
		ms.tail = nil
		return
	}
	ms.tail = ms.tail.Prev
	ms.tail.Next = nil
}

func (ms *MinStack) Top() int {
	return ms.tail.Value
}

func (ms *MinStack) GetMin() int {
	return ms.tail.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
