package heap

import (
	"errors"
	"fmt"
)

type Number interface {
	int | int32 | int64 | float32 | float64 | uint | uint32 | uint64
}

type Heap[T Number] struct {
	elements     []T
	lastAvaiable int
}

func New[T Number]() *Heap[T] {
	return &Heap[T]{
		elements:     make([]T, 1, 1),
		lastAvaiable: 0,
	}
}

func (h Heap[T]) Size() int {
	return len(h.elements)
}

func (h *Heap[T]) Insert(value T) {
	if h.lastAvaiable >= cap(h.elements) {
		newElements := make([]T, h.Size()*2)

		for i, element := range h.elements {
			newElements[i] = element
		}

		h.elements = newElements
	}

	h.elements[h.lastAvaiable] = value
	h.lastAvaiable += 1

	h.minify()
}

// Take gets and deletes the minimum value from the list.
// It returns the value
func (h *Heap[T]) Take() (T, error) {
	if h.Size() == 0 {
		return *new(T), errors.New("Heap is empty")
	}

	min := h.elements[0]
	h.elements[0] = h.elements[h.Size()-1]
	h.elements = h.elements[:h.Size()-1]

	h.minify()

	return min, nil
}

func (h *Heap[T]) minify() {
	for i := h.Size()/2 - 1; i >= 0; i-- {
		h.minifyDown(i)
	}
}

func (h *Heap[T]) minifyDown(i int) {
	min := i
	left := 2*i + 1
	right := 2*i + 2

	if left < h.Size() && h.elements[left] < h.elements[min] {
		min = left
	}

	if right < h.Size() && h.elements[right] < h.elements[min] {
		min = right
	}

	if min != i {
		h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
		h.minifyDown(min)
	}
}

func (h Heap[T]) String() string {
	return fmt.Sprintf("%v", h.elements)
}
