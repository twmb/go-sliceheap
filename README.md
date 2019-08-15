go-sliceheap
============

Package sliceheap returns a quick heap given a pointer to a slice and a less
function (akin to sort.Slice for sorting slices).

This package is for that rare time when you need a heap and do not want to make
an arbitrary type to provide `Push` and `Pop`.

Documentation
-------------

[![GoDoc](https://godoc.org/github.com/twmb/go-sliceheap?status.svg)](https://godoc.org/github.com/twmb/go-sliceheap)

Example
-------

```go
// we can create heap in one function and return it.
inner := func() sliceheap.Heap {
	a := []int{3, 2, 4, 5, 1, 0, 6}
	h := sliceheap.On(&a, func(i, j int) bool { return a[i] > a[j] })
        heap.Init(h)
        return h
}
h := inner()
// We can see the heap sort itself by checking the backing slice.
fmt.Println(h.View().([]int)) // prints [6 5 4 2 1 0 3]

heap.Push(h, 8) // push a few more elements...

// If we want to observe pushes and pops from the slice, we must save a
// pointer to the slice. This is only necessary if working on a slice
// that you did not pass directly to sliceheap.On, i.e., if you have lost
// the original pointer you used.
ptr := h.Pointer().(*[]int)
fmt.Println(*ptr) // prints [8 6 4 5 1 0 3 2]
heap.Push(h, 7)
fmt.Println(*ptr) // prints [8 7 4 6 1 0 3 2 5]

heap.Push(h, 7)
heap.Push(h, 9)
// Pop everything off, printing as we pop largest to smallest.
for h.Len() > 0 {
	largest := heap.Pop(h).(int)
	fmt.Println(largest) // prints 9, 8, 7...
}
```

Benchmarks
----------

The code works through heavy reflect usage. As such, it has many allocations.
For normal one-off heaps, this should be a nonissue.

The following benchmark shows creating a slice of N nodes, and then constantly
pushing a new largest element onto it and popping the smallest. The new largest
element constantly sifts to the bottom, then smallest sifts out to be popped.

On go1.12.7,

```
BenchmarkPushPop/1_nodes-4         	 2000000	       669 ns/op	     124 B/op	       6 allocs/op
BenchmarkPushPop/2_nodes-4         	 2000000	       694 ns/op	     128 B/op	       6 allocs/op
BenchmarkPushPop/4_nodes-4         	 2000000	       971 ns/op	     138 B/op	       8 allocs/op
BenchmarkPushPop/8_nodes-4         	 2000000	      1102 ns/op	     149 B/op	       9 allocs/op
BenchmarkPushPop/16_nodes-4        	 1000000	      1004 ns/op	     145 B/op	       9 allocs/op
BenchmarkPushPop/32_nodes-4        	 1000000	      1161 ns/op	     153 B/op	      10 allocs/op
BenchmarkPushPop/64_nodes-4        	 1000000	      1269 ns/op	     160 B/op	      11 allocs/op
BenchmarkPushPop/128_nodes-4       	 1000000	      1453 ns/op	     168 B/op	      12 allocs/op
BenchmarkPushPop/256_nodes-4       	 1000000	      1507 ns/op	     176 B/op	      13 allocs/op
BenchmarkPushPop/512_nodes-4       	 1000000	      1667 ns/op	     184 B/op	      14 allocs/op
BenchmarkPushPop/1024_nodes-4      	 1000000	      1781 ns/op	     192 B/op	      15 allocs/op
```
