// Package sliceheap returns a quick heap interface implementation given a
// pointer to a slice and a less function (akin to sort.Slice for sorting
// slices).
//
// This package is for that rare time when you need a heap and do not want to
// make an arbitrary type to provide Push and Pop.
package sliceheap

// Heap is a heap on a slice.
type Heap[T any] struct {
	slice *[]T
	less  func(i, j int) bool
}

// On returns a heap on a pointer to a slice.
//
// The heap is not initialized before returning.
func On[T any](slice *[]T, less func(i, j int) bool) Heap[T] {
	h := Heap[T]{
		slice: slice,
		less:  less,
	}
	return h
}

// View returns the backing slice the heap is on.
//
// Note that this slice is invalidated after any Push or Pop call, thus, it is
// only a view of the current slice.
func (h Heap[T]) View() []T {
	return *h.slice
}

// Pointer returns a pointer to the backing slice the heap is on.
//
// Changes to the heap can be seen by dereferencing the pointer.
func (h Heap[T]) Pointer() *[]T {
	return h.slice
}

// Swap swaps two elements in the slice.
func (h Heap[T]) Swap(i, j int) {
	v := *h.slice
	v[i], v[j] = v[j], v[i]
}

// Len returns the current length of the slice.
func (h Heap[T]) Len() int {
	return len(*h.slice)
}

// Less returns whether the element at i is less than the element at j.
func (h Heap[T]) Less(i, j int) bool {
	return h.less(i, j)
}

// Push pushes a new element onto the heap's backing slice.
func (h Heap[T]) Push(x any) {
	*h.slice = append(*h.slice, x.(T))
}

// Pop pops the smallest element off of the slice and returns it.
func (h Heap[T]) Pop() any {
	v := *h.slice
	x := v[len(v)-1]
	*h.slice = v[:len(v)-1]
	return x
}
