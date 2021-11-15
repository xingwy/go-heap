package goheap

import (
	"sort"
)

// base
type GoHeap struct {
	_pool      []*Container
	_compartor Compartor
}

// 支持sort
func (h GoHeap) Len() int { return len(h._pool) }
func (h GoHeap) Swap(i, j int) {
	h._pool[i].data, h._pool[j].data = h._pool[j].data, h._pool[i].data
}
func (h GoHeap) Less(i, j int) bool {
	return h._compartor(h._pool[i].data, h._pool[j].data) < 0
}

func NewHeap(c Compartor) *GoHeap {
	return &GoHeap{_pool: make([]*Container, 0), _compartor: c}
}

func CreateHeap(c Compartor, data []T) *GoHeap {
	if len(data) > 0 {
		_pool := make([]*Container, len(data))
		for i := 0; i < len(data); i++ {
			_pool[i] = &Container{__pointer: i, data: data[i]}
		}
		h := &GoHeap{_pool: _pool, _compartor: c}
		sort.Sort(h)

		return h
	}
	return NewHeap(c)
}

// 弹出堆顶值
func (h *GoHeap) Pop() *Container {
	if h.Len() <= 0 {
		return nil
	}

	// 调整堆
	r := h._pool[0]
	v := h._pool[h.Len()-1]
	h._pool = h._pool[:h.Len()-1]

	// 调整堆
	h.shiftdown(0, v)
	return r
}

// 堆顶值
func (h *GoHeap) Top() *Container {
	if h.Len() <= 0 {
		return &Container{data: nil, __pointer: -1}
	}
	return h._pool[0]
}

// 清除堆
func (h *GoHeap) Clear() {
	h._pool = h._pool[0:0]
}

/**
 * 新增数据
 * @param v
 */
func (h *GoHeap) Add(v T) {
	// 使用容器包裹
	container := &Container{__pointer: h.Len(), data: v}
	h._pool = append(h._pool, container)
	h.shiftup(h.Len()-1, container)
}

// 获取所有数据
func (h *GoHeap) GetPool() []*Container {
	return h._pool
}

// 获取所有数据
func (h *GoHeap) GetData() []T {
	list := make([]T, h.Len())
	for _, v := range h._pool {
		list = append(list, v)
	}
	return list
}

/**
 * 移除数据
 * @param v
 */
func (h *GoHeap) Remove(c Container) {
	pointer := c.__pointer
	if pointer >= h.Len() {
		return
	}
	if h._compartor(h._pool[pointer].Value(), c.Value()) != 0 {
		return
	}
	r := h.Pop()

	if pointer < h.Len() {
		// 重整堆
		h.shiftdown(pointer, r)
	}
}

// /**
//  * 更新数据
//  * @param v
//  */
// func Update(v T) {

// }

/**
 * 比较器
 * @param v 当前值
 * @param t 比较值
 */
func (h *GoHeap) compartor(v T, t T) bool {
	return h._compartor(v, t) < 0
}

func (h *GoHeap) shiftup(_p int, _c *Container) {
	// _p 当前节点
	f := (_p - 1) >> 1
	var v T

	for {
		if 0 >= _p {
			break
		}
		// 父容器
		v = h._pool[f].data
		if h.compartor(_c.data, v) {
			// 满足堆平衡 退出
			break
		}

		// 递归父节点
		_c.__pointer = _p
		h._pool[_p].data = v
		_p = f
		f = (_p - 1) >> 1
	}
	_c.__pointer = _p
	h._pool[_p].data = _c.data
}

/**
 * 目标值替换，元素下调
 * @param point
 * @param v
 */
func (h *GoHeap) shiftdown(_p int, _c *Container) {
	l := (_p << 1) + 1
	r := l + 1
	c := h.Len()
	var v *Container

	for {
		if l >= c {
			break
		}

		// 左分支 取左右较优分支
		v = h._pool[l]
		if r < c && h.compartor(h._pool[r].Value(), v.Value()) {
			v = h._pool[r]
			l = r
		}

		if h.compartor(_c.Value(), v.Value()) {
			break
		}
		// 迭代
		v.__pointer = _p
		h._pool[_p].data = v
		_p = l
		l = (_p << 1) + 1
		r = l + 1
	}
	// 更新目标节点指针
	_c.__pointer = _p
	h._pool[_p].data = _c

}
