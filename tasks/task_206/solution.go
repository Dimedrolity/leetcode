// Package task_206
// 206. Reverse Linked List
package task_206

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList
// Мое решение.
// Сложность по времени O(N) (один проход по списку), по памяти O(N).
// Идея в том, чтобы создавать новую элемент, который будет указывать на предыдущий.
// Не меняет старый список.
func reverseList_My(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var rev = &ListNode{head.Val, nil}
	for cur.Next != nil {
		revPrev := &ListNode{cur.Next.Val, rev}
		rev = revPrev

		cur = cur.Next
	}

	return rev
}

// reverseList
// Решение из интернета
// Сложность по времени O(N) (один проход по списку), по памяти O(1).
// Идея в том, чтобы создать 2 указателя, текущий и предыдущий.
// У текущего ставить следующим предыдущий. И временная переменная для того, чтобы не потерять следующий элемент.
// Меняет старый список.
func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur

		cur = next
	}

	return prev
}
