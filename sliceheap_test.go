package sliceheap

import (
	"container/heap"
	"fmt"
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	a := []int{3, 2, 4, 5, 1, 0, 6}
	h := On(&a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for i := 7; i < 100; i++ {
		heap.Push(h, i)
	}

	{
		current := h.Slice()
		if !reflect.DeepEqual(a, current) {
			t.Error("expected deep equal")
		}
	}

	exp := 0
	for h.Len() > 0 {
		got := heap.Pop(h).(int)
		if got != exp {
			t.Errorf("got %d != exp %d", got, exp)
		}
		exp++
	}

	{
		current := h.Slice()
		if !reflect.DeepEqual(a, current) {
			t.Error("expected deep equal")
		}
	}
}

func BenchmarkPushPop(b *testing.B) {
	for nodes := 1; nodes <= 1<<10+1; nodes <<= 1 {
		var slice []int
		for i := 0; i < nodes; i++ {
			slice = append(slice, i)
		}

		h := On(&slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})

		b.Run(fmt.Sprintf("%d_nodes", nodes), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				heap.Push(h, i)
				heap.Pop(h)
			}
		})
	}
}
