package liststack

import (
	"cmp"

	"dsa/data_structures/linked_list/singly"
)

type LinkedListStack[T cmp.Ordered] struct {
	data *singly.SinglyLinkedList[T]
}

func New[T cmp.Ordered]() *LinkedListStack[T] {
	return &LinkedListStack[T]{data: singly.New[T]()}
}
