package heapq

import (
	"math/rand"
	"testing"
)

func verify(t *testing.T, i int, pq *PQ[int]) {
	t.Helper()
	n := pq.Len()
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 < n {
		if pq.less(j1, i) {
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, pq.queue[i], j1, pq.queue[j1])
			return
		}
		verify(t, j1, pq)
	}
	if j2 < n {
		if pq.less(j2, i) {
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, pq.queue[i], j1, pq.queue[j2])
			return
		}
		verify(t, j2, pq)
	}
}

func TestInit0(t *testing.T) {
	pq := NewPQ(func(x, y int) bool {
		return x < y
	})
	for i := 20; i > 0; i-- {
		pq.Push(0) // all elements are the same
	}
	pq.init()
	verify(t, 0, pq)

	for i := 1; pq.Len() > 0; i++ {
		x := pq.Pop()
		verify(t, 0, pq)
		if x != 0 {
			t.Errorf("%d.th pop got %d; want %d", i, x, 0)
		}
	}
}

func TestInit1(t *testing.T) {
	pq := NewPQ(func(x, y int) bool {
		return x < y
	})
	for i := 20; i > 0; i-- {
		pq.Push(i) // all elements are the same
	}
	pq.init()
	verify(t, 0, pq)

	for i := 1; pq.Len() > 0; i++ {
		x := pq.Pop()
		verify(t, 0, pq)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func Test(t *testing.T) {
	pq := NewPQ(func(x, y int) bool {
		return x < y
	})
	verify(t, 0, pq)

	for i := 20; i > 10; i-- {
		pq.Push(i)
	}
	pq.init()
	verify(t, 0, pq)

	for i := 10; i > 0; i-- {
		pq.Push(i)
		verify(t, 0, pq)
	}

	for i := 1; pq.Len() > 0; i++ {
		x := pq.Pop()
		if i < 20 {
			pq.Push(20 + i)
		}
		verify(t, 0, pq)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func TestRemove0(t *testing.T) {
	pq := NewPQ(func(x, y int) bool {
		return x < y
	})
	for i := 0; i < 10; i++ {
		pq.Push(i)
	}
	verify(t, 0, pq)

	for pq.Len() > 0 {
		i := pq.Len() - 1
		x := pq.Remove(i)
		if x != i {
			t.Errorf("Remove(%d) got %d; want %d", i, x, i)
		}
		verify(t, 0, pq)
	}
}

func TestRemove1(t *testing.T) {
	pq := NewPQ(func(x, y int) bool {
		return x < y
	})
	for i := 0; i < 10; i++ {
		pq.Push(i)
	}
	verify(t, 0, pq)

	for i := 0; pq.Len() > 0; i++ {
		x := pq.Remove(0)
		if x != i {
			t.Errorf("Remove(0) got %d; want %d", x, i)
		}
		verify(t, 0, pq)
	}
}

func TestRemove2(t *testing.T) {
	N := 10

	pq := NewPQ(func(x, y int) bool {
		return x < y
	})
	for i := 0; i < N; i++ {
		pq.Push(i)
	}
	verify(t, 0, pq)

	m := make(map[int]bool)
	for pq.Len() > 0 {
		m[pq.Remove((pq.Len()-1)/2)] = true
		verify(t, 0, pq)
	}

	if len(m) != N {
		t.Errorf("len(m) = %d; want %d", len(m), N)
	}
	for i := 0; i < len(m); i++ {
		if !m[i] {
			t.Errorf("m[%d] doesn't exist", i)
		}
	}
}

func TestFix(t *testing.T) {
	pq := NewPQ(func(x, y int) bool {
		return x < y
	})
	verify(t, 0, pq)

	for i := 200; i > 0; i -= 10 {
		pq.Push(i)
	}
	verify(t, 0, pq)

	if pq.queue[0] != 10 {
		t.Fatalf("Expected head to be 10, was %d", pq.queue[0])
	}
	pq.queue[0] = 210
	pq.Fix(0)
	verify(t, 0, pq)

	for i := 100; i > 0; i-- {
		elem := rand.Intn(pq.Len())
		if i&1 == 0 {
			pq.queue[elem] *= 2
		} else {
			pq.queue[elem] /= 2
		}
		pq.Fix( elem)
		verify(t, 0,pq)
	}
}
