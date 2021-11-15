package goheap

import "sort"

// base
type GoHeap struct {
	_pool      []Container
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
	return &GoHeap{_pool: make([]T, 0), _compartor: c}
}

func CreateHeap(c Compartor, data []T) *GoHeap {
	if len(data) > 0 {
		_pool := make([]T, len(data))
		for i := 0; i < len(data); i++ {
			_pool[i] = T{__pointer: i, data: data[i]}
		}
		h := &GoHeap{_pool: _pool, _compartor: c}
		sort.Sort(h)

		return h
	}
	return NewHeap(c)
}

// 弹出堆顶值
func (h *GoHeap) Pop() T {
	if h.Len() <= 0 {
		return nil
	}

	// 调整堆
	r := h._pool[0]
	v := h._pool[h.Len()-1]
	h._pool = h._pool[:h.Len()-1]

	// 调整堆
	// h.shiftdown(0, v)
	return r.data
}

// 堆顶值
func (h *GoHeap) Top() T {
	if h.Len() <= 0 {
		return nil
	}
	return h._pool[0].data
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
	container := Container{ __pointer: h.Len(), data: v}
	h._pool = append(h._pool, container)
	h.shiftup(h.Len()-1, container)
}

// // 获取所有数据
// func GetDatas(h *GoHeap) []T {
// 	// return this._rankHeap.getDatas()
// 	return h._pool
// }

// /**
//  * 移除数据
//  * @param v
//  */
// func Remove(h *GoHeap, v T) {
// 	// if (v.__rank == 0) {
// 	// 	this._rankHeap.remove(v);
// 	// 	return;
// 	// } else {
// 	// 	this._rank.splice(v.__rank - 1, 1);
// 	// }
// 	// return this.onUpdateRank();
// }

// /**
//  * 更新数据
//  * @param v
//  */
// func Update(v T) {

// }

// /**
//  * 比较器
//  * @param v 当前值
//  * @param t 比较值
//  */
// func compartor(h *GoHeap, v T, t T) int {
// 	return h._compartor(v, t)
// }

// // 更新排行
// func onUpdateRank(h *GoHeap) {
// 	// 维护排行
// 	tailIndex := this._rank.length - 1;
// 	if (tailIndex < this._rankSize - 1) {
// 		if (this._rankHeap.size() == 0) {
// 			return;
// 		}

// 		let top = this._rankHeap.pop();
// 		top.__rank = tailIndex + 2;
// 	} else {
// 		if (this._rankHeap.size() == 0) {
// 			return;
// 		}
// 		let tail = this._rank[tailIndex];
// 		let top = this._rankHeap.pop();
// 		if (this.compartor(tail, top)) {
// 			this._rank[tailIndex] = top;
// 			this._rank.sort((l: T, r: T): number => this.compartor(l, r) ? -1 : 1)
// 			this._rank.forEach((v: T, index: number): void => {
// 				v.__rank = index + 1;
// 			});
// 			tail.__rank = 0;
// 			this._rankHeap.push(tail);
// 		}
// 	}
// }

func (h *GoHeap) shiftup(_p int, _c Container) {
	f := (_p - 1) >> 1
	var v Container

	for {
		if 0 == _p {
			break
		}

		v = h._pool[f]
		if h._compartor(_c.data, v) < 0 {
			break
		}
		// 递归父节点
		_c.__pointer = _p
		h._pool[_p] = v
		_p = f
		f = (_p - 1) >> 1
	}

	_c.__pointer = _p
	h._pool[_p] = _c
}

// /**
//  * 目标值替换，元素下调
//  * @param point
//  * @param v
//  */
// func shiftdown(_p: number, _v: T): void {
// 	let l: number = (_p << 1) + 1;
// 	let r: number = l + 1;
// 	let c: number = this._datas.length;
// 	let v: T;

// 	while (l < c) {
// 		v = this._datas[l];
// 		// 获取左右节点优先值
// 		if (r < c && this.compartor(this._datas[r], v)) {
// 			v = this._datas[r];
// 			l = r;
// 		}

// 		if (this.compartor(_v, v)) {
// 			break;
// 		}

// 		// 迭代
// 		v.__pointer = _p;
// 		this._datas[_p] = v;
// 		_p = l;
// 		l = (_p << 1) + 1;
// 		r = l + 1;
// 	}

// 	// 更新目标节点指针
// 	_v.__pointer = _p;
// 	this._datas[_p] = _v;
// }
