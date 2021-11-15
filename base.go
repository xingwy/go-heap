package goheap

// 类型定义

type T interface{}

type Container struct {
	__pointer int
	data      T
}

func (c Container) Value() T {
	return c.data
}

// type  []T
type Compartor func(a, b interface{}) int
