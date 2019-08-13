package sliceheap

import (
	"container/heap"
	"fmt"
	"reflect"
	"testing"
)

func TestOnAll(t *testing.T) {
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

func TestPushKeys(t *testing.T) {
	var s []int
	h := On(&s, func(i, j int) bool {
		return s[i] < s[j]
	})

	m := map[int]struct{}{
		9: struct{}{},
		8: struct{}{},
		7: struct{}{},
		6: struct{}{},
		5: struct{}{},
		4: struct{}{},
		3: struct{}{},
		2: struct{}{},
		1: struct{}{},
		0: struct{}{},
	}
	h.PushKeys(m)

	for exp := 0; exp < 10; exp++ {
		got := heap.Pop(h).(int)
		if got != exp {
			t.Errorf("got %d != exp %d", got, exp)
		}
	}
	if h.Len() != 0 {
		t.Errorf("expected no more entries")
	}
}

func TestPushValues(t *testing.T) {
	for _, i := range []interface{}{
		map[string]int{
			"a": 0,
			"b": 1,
			"c": 2,
			"d": 3,
			"e": 4,
		},
		[5]int{4, 3, 2, 1, 0},
		[]int{4, 3, 2, 1, 0},
	} {
		var s []int
		h := On(&s, func(i, j int) bool {
			return s[i] < s[j]
		})

		h.PushValues(i)
		for exp := 0; exp < 5; exp++ {
			got := heap.Pop(h).(int)
			if got != exp {
				t.Errorf("got %d != exp %d", got, exp)
			}
		}
		if h.Len() != 0 {
			t.Errorf("expected no more entries")
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

func BenchmarkPushValues(b *testing.B) {
	m := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
	}
	for i := 0; i < b.N; i++ {
		var s []int
		h := On(&s, func(i, j int) bool {
			return s[i] < s[j]
		})
		h.PushValues(m)
	}
}
