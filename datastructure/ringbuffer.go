package datastructure

type RingBuffer[T any] struct {
	data   []T
	size   int
	rhead  int
	whead  int
	isFull bool
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		data:   make([]T, size),
		size:   size,
		rhead:  0,
		whead:  0,
		isFull: false,
	}
}

func (r *RingBuffer[T]) Clear() {
	r.rhead = 0
	r.whead = 0
	r.data = nil
	r.data = make([]T, r.size)
}

func (r *RingBuffer[T]) Read(count int) []T {
	if r.Readable() == 0 || count == 0 {
		return []T{}
	}

	readcnt := min(count, r.Readable())
	rst := make([]T, readcnt)

	if r.rhead+readcnt >= r.size {
		remainUnilEnd := r.size - r.rhead
		copy(rst, r.data[r.rhead:])
		r.rhead = 0

		remain := readcnt - remainUnilEnd
		copy(rst[remainUnilEnd:], r.data[:remain])
		r.rhead += remain
	} else {
		copy(rst, r.data[r.rhead:r.rhead+readcnt])
		r.rhead += readcnt
	}
	r.isFull = false

	return rst
}

func (r *RingBuffer[T]) Write(data []T) int {
	if len(data) == 0 || r.Writable() == 0 {
		return 0
	}

	var writed int
	if r.whead >= r.rhead {
		writableToEnd := r.size - r.whead
		writed = min(writableToEnd, len(data))
	} else {
		writed = min(r.Writable(), len(data))
	}
	// fmt.Printf("whead#1:%d, %d, %d \n", r.whead, writed, r.Writable())

	copy(r.data[r.whead:], data[:writed])
	r.whead = (r.whead + writed) % r.size

	// fmt.Printf("whead#2:%d, %d, %d \n", r.whead, writed, r.Writable())

	// isFull?
	if writed > 0 && r.whead == r.rhead {
		r.isFull = true
	}
	remain := len(data) - writed
	if remain > 0 && r.Writable() > 0 {
		writed += r.Write(data[writed:])
	}

	return writed
}

func (r *RingBuffer[T]) Readable() int {
	if r.isFull {
		return r.size
	}

	if r.whead < r.rhead {
		return (r.size - r.rhead + r.whead)
	}

	return r.whead - r.rhead
}

func (r *RingBuffer[T]) Writable() int {
	// fmt.Printf("whead:%d\n", r.whead)
	return r.size - r.Readable()
}

// func main() {
// 	var rBuff *RingBuffer[byte]
// 	rBuff = NewRingBuffer[byte](10)

// 	fmt.Printf("#### writable#1 :%d \n", rBuff.Writable())
// 	rBuff.Write([]byte{1, 2, 3, 4, 5, 6, 7})
// 	fmt.Printf("#### readable#1 :%d \n", rBuff.Readable())
// 	rdt := rBuff.Read(3)
// 	fmt.Printf("#### read#1 :%v \n", rdt)
// 	fmt.Printf("#### writable#1-1 :%d \n", rBuff.Writable())

// 	rdt12 := rBuff.Read(4)
// 	fmt.Printf("#### read#1-1 :%v \n", rdt12)
// 	fmt.Printf("#### writable#1-1 :%d \n", rBuff.Writable())
// 	fmt.Printf("--------------------------------\n")

// 	rBuff.Write([]byte{11, 12, 13, 14, 15})
// 	fmt.Printf("#### writable#2 :%d \n", rBuff.Writable())
// 	fmt.Printf("#### readable#2 :%d \n", rBuff.Readable())
// 	rdt2 := rBuff.Read(rBuff.Readable())
// 	fmt.Printf("#### read#2 :%v \n", rdt2)

// }
