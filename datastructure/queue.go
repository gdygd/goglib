package datastructure

import "sync"

var qmtx = &sync.Mutex{} // comm manage mutex

type SlQueye[T any] struct {
	arr []T
}

func NewSlQueue[T any]() *SlQueye[T] {
	return &SlQueye[T]{arr: []T{}}
}

func (q *SlQueye[T]) Push(val T) {
	qmtx.Lock()
	q.arr = append(q.arr, val)
	qmtx.Unlock()
}

func (q *SlQueye[T]) Pop() (T, bool) {

	var front T
	if len(q.arr) == 0 {
		return front, false
	}

	qmtx.Lock()
	front = q.arr[0]
	q.arr = q.arr[1:]
	qmtx.Unlock()
	return front, true
}

func (q *SlQueye[T]) QClear() {
	qmtx.Lock()
	q.arr = []T{}
	qmtx.Unlock()
}
