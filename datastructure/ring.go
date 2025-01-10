package datastructure

type Ring struct {
	data []int
	size int
	head int
}

func NewRing(size int) *Ring {
	return &Ring{
		data: make([]int, size),
		size: size,
		head: 0,
	}
}

func (r *Ring) Add(value int) {
	r.data[r.head] = value
	r.head = (r.head + 1) % r.size
}

func (r *Ring) Get(index int) int {
	if index < 0 || index >= r.size {
		panic("index out of range")
	}
	return r.data[(r.head+index)%r.size]
}

func (r *Ring) Move(pos int) {
	if r.head+pos > 0 {
		r.head = (r.head + pos) % r.size
	} else {
		pos2 := pos % (r.size * -1)
		r.head = r.head + pos2 + r.size
	}
}

// func main() {
// 	ring1 := NewRing(7)
// 	ring2 := NewRing(7)

// 	for i := 1; i <= 7; i++ {
// 		ring1.Add(i)
// 		ring2.Add(i)
// 	}

// 	for i := 0; i < ring1.size; i++ {
// 		fmt.Printf("Element at %d: %d\n", i, ring1.Get(i))
// 	}

// 	ring1.Move(2)
// 	ring2.Move(0)
// 	fmt.Printf("head:(%d, %d) \n", ring1.head, ring2.head)
// 	for i := 0; i < ring1.size; i++ {
// 		fmt.Printf("Ring1 Element at %d: %d\n", i, ring1.Get(i))
// 	}
// 	for i := 0; i < ring2.size; i++ {
// 		fmt.Printf("Ring 2Element at %d: %d\n", i, ring2.Get(i))
// 	}
// }
