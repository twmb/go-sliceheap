package sliceheap_test

import (
	"container/heap"
	"fmt"

	"github.com/twmb/go-sliceheap"
)

func ExampleHeap() {
	// inner shows that we can create heap in one function and return it.
	inner := func() sliceheap.Heap[int] {
		a := []int{3, 2, 4, 5, 1, 0, 6}
		h := sliceheap.On(&a, func(i, j int) bool { return a[i] > a[j] })
		heap.Init(h)
		return h
	}
	h := inner()
	// We can see the heap sort itself by checking the backing slice.
	fmt.Println(h.View())
	// Push a few more elements.
	heap.Push(h, 8)

	// If we want to observe pushes and pops from the slice, we must save a
	// pointer to the slice. This is only necessary if working on a slice
	// that you did not pass directly to sliceheap.On, i.e., you have lost
	// the original pointer you used.
	ptr := h.Pointer()
	fmt.Println(*ptr)
	heap.Push(h, 7)
	fmt.Println(*ptr)

	heap.Push(h, 9)
	// Pop everything off, printing as we pop largest to smallest.
	for h.Len() > 0 {
		largest := heap.Pop(h).(int)
		fmt.Println(largest)
	}
	// Output:
	// [6 5 4 2 1 0 3]
	// [8 6 4 5 1 0 3 2]
	// [8 7 4 6 1 0 3 2 5]
	// 9
	// 8
	// 7
	// 6
	// 5
	// 4
	// 3
	// 2
	// 1
	// 0
}

func ExampleOn() {
	a := []int{3, 2, 4, 5, 1, 0, 6}
	h := sliceheap.On(&a, func(i, j int) bool { return a[i] > a[j] })
	heap.Init(h)
	fmt.Println(heap.Pop(h).(int))
	// Output:
	// 6
}
