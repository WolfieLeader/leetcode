package arrayqueue

import (
	"cmp"

	"dsa/data_structures/array/dynamic"
)

type ArrayQueue[T cmp.Ordered] struct {
	data *dynamic.DynamicArray[T]
}

func New[T cmp.Ordered]() *ArrayQueue[T] {
	return &ArrayQueue[T]{data: dynamic.New[T]()}
}
