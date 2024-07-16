package linked_list

import "cmp"

type ListNode[T cmp.Ordered] struct {
	Data T
	Next *ListNode[T]
}

func RemoveNode[T cmp.Ordered](data T, head *ListNode[T]) *ListNode[T] {
	slow, fast := head, head.Next
	if slow.Data == data {
		slow.Next = nil
		return fast
	}
	for fast != nil {
		if fast.Data == data {
			slow.Next = fast.Next
			return head
		}
		fast = fast.Next
		slow = slow.Next
	}
	return head
}

func ReverseLinkedList[T cmp.Ordered](head *ListNode[T]) *ListNode[T] {
	slow, fast := head, head.Next
	for fast != nil {
		slow.Next = nil
		slow = fast
		fast = fast.Next
	}
	return slow
}
