package heap_test

import (
	"fmt"
	"testing"

	"github.com/cgault/go-data-structures/heap"
)

func Test(t *testing.T) {
	m := &heap.MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
