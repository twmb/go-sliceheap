package sliceheap_test

import (
	"container/heap"
	"fmt"

	"github.com/twmb/go-sliceheap"
)

func ExampleHeap() {
	// inner shows that we can create heap in one function and return it.
	inner := func() sliceheap.Heap {
		a := []int{3, 2, 4, 5, 1, 0, 6}
		return sliceheap.On(&a, func(i, j int) bool { return a[i] > a[j] })
	}
	h := inner()
	// We can see the heap sort itself by checking the backing slice.
	fmt.Println(h.Slice().([]int))
	// Push a few more elements.
	heap.Push(h, 8)
	heap.Push(h, 7)
	heap.Push(h, 9)
	// Pop everything off, printing as we pop largest to smallest.
	for h.Len() > 0 {
		largest := heap.Pop(h).(int)
		fmt.Println(largest)
	}
	// Output:
	// [6 5 4 2 1 0 3]
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
