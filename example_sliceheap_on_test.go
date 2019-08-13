package sliceheap_test

import (
	"fmt"

	"github.com/twmb/go-sliceheap"
)

func ExampleOn() {
	a := []int{3, 2, 4, 5, 1, 0, 6}
	h := sliceheap.On(&a, func(i, j int) bool { return a[i] > a[j] })
	fmt.Println(h.Len())
	// Output:
	// 7
}
