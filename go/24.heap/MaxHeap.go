package main

import "fmt"

type MaxHeap struct {
	data []int
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func left(idx int) int {
	return idx*2 + 1
}

func right(idx int) int {
	return idx*2 + 2
}

func (h MaxHeap) String() string {
	return fmt.Sprintf("MaxHeap: %v", h.data)
}

func (m *MaxHeap) insert(d int) {
	newData := append(m.data, d)

	for idx := len(newData) - 1; idx > 0; {
		p := parent(idx)
		if newData[idx] <= newData[p] {
			break
		}
		newData[idx], newData[p] = newData[p], newData[idx]
		idx = p
	}

	m.data = newData
}

func (h *MaxHeap) pop() (int, bool) {
	size := len(h.data)
	if size == 0 {
		return -1, false
	}
	top := h.data[0]
	if size == 1 {
		h.data = []int{}
		return top, true
	}

	size--
	h.data[0] = h.data[size]
	h.data = h.data[:size]

	idx := 0
	for {
		l, r := left(idx), right(idx)
		next := l
		if r < size {
			if h.data[l] < h.data[r] {
				next = r
			}
		} else if l >= size {
			break
		}

		if h.data[idx] < h.data[next] {
			h.data[idx], h.data[next] = h.data[next], h.data[idx]
			idx = next
		} else {
			break
		}
	}

	return top, true
}

func maxHeapMain() {
	h := MaxHeap{[]int{50, 40, 25, 20, 35, 10, 15}}

	fmt.Println("Initial", h)

	h.insert(45)
	fmt.Println("Inserting 45 to heap")
	fmt.Println(h)
	fmt.Println("Updated", h)

	h.insert(75)
	fmt.Println("Inserting 75 to heap")
	fmt.Println("Updated", h)

	top, exists := h.pop()
	for exists {
		fmt.Println("Took top", top)
		fmt.Println("Updated", h)
		top, exists = h.pop()
	}
}
