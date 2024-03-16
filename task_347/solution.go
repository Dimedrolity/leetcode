// Package task_347
// 347. Top K Frequent Elements
package task_347

import (
	"container/heap"
	"slices"
	"sort"
)

// topKFrequent
// Мое решение.
// Сложность по времени O(N*K), по памяти O(N). Если K=N, то сложность по времени O(N^2).
// Идея в том, чтобы создать частотный словарь и взять из него k элементов с максимальной частотой (макс значением словаря).
// Решение должно быть лучше O(N*logN), т.е. нельзя использовать сортировку.
func topKFrequent_N2(nums []int, k int) []int {
	wl := wordList(nums)
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		topFrequent := findTopFrequent(wl)
		result = append(result, topFrequent)
		delete(wl, topFrequent)
	}
	return result
}

func wordList(nums []int) map[int]int {
	wl := make(map[int]int)
	for _, num := range nums {
		wl[num] += 1
	}
	return wl
}

func findTopFrequent(wl map[int]int) int {
	values := make([]int, 0, len(wl))
	for _, value := range wl {
		values = append(values, value)
	}
	maxValue := slices.Max(values)
	for value, freq := range wl {
		if freq == maxValue {
			return value
		}
	}

	return 0
}

// An NumbersMinHeap is a min-heap of FrequencyNum.
type NumbersMinHeap []FrequencyNum

func (h NumbersMinHeap) Len() int           { return len(h) }
func (h NumbersMinHeap) Less(i, j int) bool { return h[i].Frequency < h[j].Frequency }
func (h NumbersMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NumbersMinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(FrequencyNum))
}

func (h *NumbersMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type FrequencyNum struct {
	Number    int
	Frequency int
}

// topKFrequent
// Мое решение.
// Сложность по времени O(N*logN), по памяти O(N).
// Идея в том, чтобы создать частотный словарь и построить Min Heap на его основе.
// Куча отсортирована в порядке зависящем от частоты встречаемости, таким образом в корне будет минимально встречающийся элемент.
// Операции над кучей за O(logN).
func topKFrequent_Heap(nums []int, k int) []int {
	wl := wordList(nums)
	freqNums := make([]FrequencyNum, 0)
	for value, freq := range wl {
		freqNums = append(freqNums, FrequencyNum{value, freq})
	}
	h := &NumbersMinHeap{}
	for _, freqNum := range freqNums {
		if h.Len() < k {
			heap.Push(h, freqNum)
		} else {
			minFreq := heap.Pop(h).(FrequencyNum)

			if minFreq.Frequency > freqNum.Frequency {
				heap.Push(h, minFreq)
			} else {
				heap.Push(h, freqNum)
			}
		}
	}
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, h.Pop().(FrequencyNum).Number)
	}

	return result
}

// topKFrequent
// Мое решение.
// Сложность по времени O(N*logN), по памяти O(N).
// Идея в том, чтобы создать частотный словарь, отсортировать его по значениям и вернуть значения.
func topKFrequent_Sort(nums []int, k int) []int {
	wl := wordList(nums)
	freqNums := make([]FrequencyNum, 0)
	for value, freq := range wl {
		freqNums = append(freqNums, FrequencyNum{value, freq})
	}
	sort.Slice(freqNums, func(i, j int) bool { return freqNums[i].Frequency > freqNums[j].Frequency })

	values := make([]int, 0)
	for _, freqNum := range freqNums {
		values = append(values, freqNum.Number)
	}
	return values[:k]
}

// An NumbersMaxHeap is a min-heap of FrequencyNum.
type NumbersMaxHeap []FrequencyNum

func (h NumbersMaxHeap) Len() int           { return len(h) }
func (h NumbersMaxHeap) Less(i, j int) bool { return h[i].Frequency > h[j].Frequency }
func (h NumbersMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NumbersMaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(FrequencyNum))
}

func (h *NumbersMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// topKFrequent
// Мое решение, но с подсказкой, что heap.Init зарботает за O(N).
// Сложность по времени O(k*logN), по памяти O(N).
// Идея в том, чтобы создать частотный словарь, положить все элементы в Max Heap за O(N) и взять k элементов.
// Важно, что heap.Init зарботает за O(N) !!!
func topKFrequent(nums []int, k int) []int {
	wl := wordList(nums)

	h := make(NumbersMaxHeap, 0)
	for value, freq := range wl {
		h = append(h, FrequencyNum{value, freq})
	}

	heap.Init(&h)

	values := make([]int, 0)
	for i := 0; i < k; i++ {
		values = append(values, heap.Pop(&h).(FrequencyNum).Number)
	}

	return values[:k]
}
