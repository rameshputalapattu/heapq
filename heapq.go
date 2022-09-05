package heapq

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type PQ[T any] struct {
	queue []T
	index int
	Less  func(x, y T) bool
}

func (pq *PQ[T]) Push(item T) {
	n := pq.Len()
	pq.index = n
	pq.queue = append(pq.queue, item)

	pq.up(pq.Len() - 1)
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func (pq *PQ[T]) init() {
	// heapify
	n := pq.Len()
	for i := n/2 - 1; i >= 0; i-- {
		pq.down(i, n)
	}
}

func (pq *PQ[T]) PrintAll() {

	fmt.Println(pq.queue)

}

func NewPQ[T any](less func(x, y T) bool) *PQ[T] {
	return &PQ[T]{
		queue: make([]T, 0),
		Less:  less,
	}
}

func NewPQWithItems[T any](items []T, less func(x, y T) bool) *PQ[T] {
	pq := &PQ[T]{
		queue: items,
		Less:  less,
	}

	pq.init()
	return pq

}

type Number interface {
	int|int32|int64|int8|float32|float64
}

func NewPQNumber[T Number]() *PQ[T] {
	pq := &PQ[T] {
		queue: make([]T, 0),
		Less: func(x,y T) bool {
			return x < y
		},
	}

	return pq
}


func NewPQOrdered[T constraints.Ordered]() *PQ[T] {
	pq := &PQ[T]{
		queue: make([]T, 0),
		Less: func(x, y T) bool {
			switch any(x).(type) {
			case int:

				return any(x).(int) < any(y).(int)
			case int8:
				return any(x).(int8) < any(y).(int8)
			case int16:
				return any(x).(int16) < any(y).(int16)
			case int32:
				return any(x).(int32) < any(y).(int32)
			case int64:
				return any(x).(int32) < any(y).(int32)
			case float32:
				return any(x).(float32) < any(y).(float32)
			case float64:
				return any(x).(float64) < any(y).(float64)
	
			
			

			default:
				return false

			}
		},
	}

	return pq
}

func (pq *PQ[T]) less(i, j int) bool {
	return pq.Less(pq.queue[i], pq.queue[j])
}

func (pq *PQ[T]) swap(i, j int) {
	pq.queue[j], pq.queue[i] = pq.queue[i], pq.queue[j]

}

func (pq *PQ[T]) Len() int {
	return len(pq.queue)

}

func (pq *PQ[T]) Fix(i int) {
	if !pq.down(i, pq.Len()) {
		pq.up(i)
	}
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (pq *PQ[T]) Remove(i int) T {
	n := pq.Len() - 1
	if n != i {
		pq.swap(i, n)
		if !pq.down(i, n) {
			pq.up(i)
		}
	}
	old := pq.queue
	item := old[n]

	pq.queue = old[0:n]
	return item
}

func (pq *PQ[T]) Pop() T {
	n := pq.Len() - 1
	pq.swap(0, n)
	pq.down(0, n)
	old := pq.queue
	item := old[n]

	pq.queue = old[0:n]

	return item
}

func (pq *PQ[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		j = i
	}
}

func (pq *PQ[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && pq.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		i = j
	}
	return i > i0
}
