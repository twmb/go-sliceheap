// Package sliceheap returns a quick heap interface implementation given a
// pointer to a slice and a less function (akin to sort.Slice for sorting
// slices).
//
// This package is for that rare time when you need a heap and do not want to
// make an arbitrary type to provide Push and Pop.
package sliceheap

import "reflect"

// Heap is a heap on a slice.
type Heap struct {
	slicePtr reflect.Value
	less     func(i, j int) bool
}

// On returns a heap on a pointer to a slice.
//
// The heap is not initialized before returning.
func On(slicePtr interface{}, less func(i, j int) bool) Heap {
	h := Heap{
		slicePtr: reflect.ValueOf(slicePtr),
		less:     less,
	}
	return h
}

// View returns the backing slice the heap is on.
//
// Note that this slice is invalidated after any Push or Pop call, thus, it is
// only a view of the current slice.
func (h Heap) View() interface{} {
	return reflect.Indirect(h.slicePtr).Interface()
}

// Pointer returns a pointer to the backing slice the heap is on.
//
// Changes to the heap can be seen by dereferencing the pointer.
func (h Heap) Pointer() interface{} {
	return h.slicePtr.Interface()
}

// Peek returns the element at index 0 in the heap, corresponding to the
// current smallest value.
func (h Heap) Peek() interface{} {
	return reflect.Indirect(h.slicePtr).Index(0).Interface()
}

// Swap swaps two elements in the slice.
func (h Heap) Swap(i, j int) {
	slice := reflect.Indirect(h.slicePtr)
	l := slice.Index(i)
	m := l.Interface() // copy out our value; the temporary middle
	r := slice.Index(j)
	l.Set(r)
	r.Set(reflect.ValueOf(m))
}

// Len returns the current length of the slice.
func (h Heap) Len() int {
	return reflect.Indirect(h.slicePtr).Len()
}

// Less returns whether the element at i is less than the element at j.
func (h Heap) Less(i, j int) bool {
	return h.less(i, j)
}

// Push pushes a new element onto the heap's backing slice.
func (h Heap) Push(x interface{}) {
	slice := reflect.Indirect(h.slicePtr)
	slice.Set(reflect.Append(slice, reflect.ValueOf(x)))
}

// Pop pops the smallest element off of the slice and returns it.
func (h Heap) Pop() interface{} {
	slice := reflect.Indirect(h.slicePtr)
	len := slice.Len()
	last := slice.Index(len - 1)
	slice.SetLen(len - 1)
	return last.Interface()
}
