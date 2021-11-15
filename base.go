package goheap

// 类型定义

type T interface{}

type Container struct {
	__pointer int
	data      T
}

// type  []T
type Compartor func(a, b interface{}) int
