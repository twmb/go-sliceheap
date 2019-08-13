// Package sliceheap returns a quick heap given a pointer to a slice and a less
// function (akin to sort.Slice for sorting slices).
//
// This package is for that rare time when you need a heap and do not want to
// make an arbitrary type to provide Push and Pop.
//
// To further aid in making life easier, the heap can be quickly populated from
// a map, slice, or array with the PushKeys or PushValues functions.
package sliceheap

import (
	"container/heap"
	"reflect"
)

// Heap is a heap on a slice.
type Heap struct {
	slicePtr reflect.Value
	less     func(i, j int) bool
}

// On returns a heap on a pointer to a slice, initializing the heap before
// returning.
func On(slicePtr interface{}, less func(i, j int) bool) Heap {
	h := Heap{
		slicePtr: reflect.ValueOf(slicePtr),
		less:     less,
	}
	heap.Init(h)
	return h
}

// Slice returns the backing slice the heap is on.
func (h Heap) Slice() interface{} {
	return reflect.Indirect(h.slicePtr).Interface()
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

// PushKeys pushes all keys of a map into the heap.
//
// This can be used for quickly populating the heap from a map.
func (h Heap) PushKeys(m interface{}) {
	iter := reflect.ValueOf(m).MapRange()
	for iter.Next() {
		heap.Push(h, iter.Key().Interface())
	}
}

// PushValues pushes all values of either a map, array, or slice into the
// heap.
//
// This can be used for quickly populating the heap from a map or from
// a slice / array that you do not want to modify.
func (h Heap) PushValues(t interface{}) {
	tv := reflect.ValueOf(t)
	switch tv.Kind() {
	case reflect.Map:
		iter := tv.MapRange()
		for iter.Next() {
			heap.Push(h, iter.Value().Interface())
		}
	case reflect.Slice, reflect.Array:
		for i := tv.Len() - 1; i >= 0; i-- {
			heap.Push(h, tv.Index(i).Interface())
		}
	default:
		panic("unsupported push kind " + tv.Kind().String())
	}
}
