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
inner := func() sliceheap.Heap[int] {
	a := []int{3, 2, 4, 5, 1, 0, 6}
	h := sliceheap.On(&a, func(i, j int) bool { return a[i] > a[j] })
	heap.Init(h)
	return h
}
h := inner()
// We can see the heap sort itself by checking the backing slice.
fmt.Println(h.View()) // prints [6 5 4 2 1 0 3]

heap.Push(h, 8) // push a few more elements...

// If we want to observe pushes and pops from the slice, we must save a
// pointer to the slice. This is only necessary if working on a slice
// that you did not pass directly to sliceheap.On, i.e., if you have lost
// the original pointer you used.
ptr := h.Pointer()
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

The following benchmark shows creating a slice of N nodes, and then constantly
pushing a new largest element onto it and popping the smallest. The new largest
element constantly sifts to the bottom, then smallest sifts out to be popped.

On go1.20.2

```
BenchmarkPushPop/1_nodes-8         	18511695	        65.84 ns/op	      47 B/op	       3 allocs/op
BenchmarkPushPop/2_nodes-8         	17404094	        68.47 ns/op	      47 B/op	       3 allocs/op
BenchmarkPushPop/4_nodes-8         	15398985	        77.69 ns/op	      48 B/op	       3 allocs/op
BenchmarkPushPop/8_nodes-8         	13573598	        87.86 ns/op	      47 B/op	       3 allocs/op
BenchmarkPushPop/16_nodes-8        	11884834	        99.30 ns/op	      47 B/op	       3 allocs/op
BenchmarkPushPop/32_nodes-8        	10706637	       111.5 ns/op	      47 B/op	       3 allocs/op
BenchmarkPushPop/64_nodes-8        	 9599702	       125.0 ns/op	      47 B/op	       3 allocs/op
BenchmarkPushPop/128_nodes-8       	 8897252	       135.6 ns/op	      47 B/op	       3 allocs/op
BenchmarkPushPop/256_nodes-8       	 8381521	       146.2 ns/op	      47 B/op	       3 allocs/op
BenchmarkPushPop/512_nodes-8       	 8005512	       153.4 ns/op	      47 B/op	       3 allocs/op
BenchmarkPushPop/1024_nodes-8      	 7690998	       162.2 ns/op	      47 B/op	       3 allocs/op
```
