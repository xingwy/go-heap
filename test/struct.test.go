package test

import (
	"fmt"

	goheap "github.com/xingwy/go-heap"
)

type People struct {
	Age  int
	Name string
}

func TestStruct() {

	var names = []string{"first", "second", "third", "fouth"}
	var age = []int{21, 23, 20, 19}
	var peopleList = []People{}
	for index, name := range names {
		peo := People{Age: age[index], Name: name}
		peopleList = append(peopleList, peo)
	}
	var heap = goheap.NewHeap(func(a, b goheap.T) int {
		var _a = a.(People)
		var _b = b.(People)
		return _b.Age - _a.Age
	})
	for _, p := range peopleList {
		heap.Add(p)
	}

	fmt.Println(heap.Top().Value())
}
