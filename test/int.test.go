package test

import (
	"fmt"
	"math/rand"
	"sort"

	goheap "github.com/xingwy/go-heap"
)

type Array []int

func (a Array) Len() int { return len(a) }
func (a Array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Array) Less(i, j int) bool {
	return a[i] > a[j]
}

func TestInt() {

	var nums = make([]goheap.T, 100)
	var valid = make(Array, 100)
	for i := 0; i < 100; i++ {
		v := rand.Intn(10000)
		nums[i] = v
		valid[i] = v
	}
	sort.Sort(valid)
	var c = func(a, b goheap.T) int {
		var _a = a.(int)
		var _b = b.(int)
		return _b - _a
	}
	var heap = goheap.CreateHeap(c, nums)
	for _, v := range valid {
		var t = heap.Pop().Value()
		if t != v {
			fmt.Println(t, v)
		}
	}

}
