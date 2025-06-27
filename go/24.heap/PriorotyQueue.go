package main

import "fmt"

type PQData struct {
	val      int
	priority int
}

func (d PQData) String() string {
	return fmt.Sprintf("(v: %d, p:%d)", d.val, d.priority)
}

func ascending(first PQData, second PQData) bool {
	return first.priority > second.priority
}

func descending(first PQData, second PQData) bool {
	return first.priority < second.priority
}

type PriorityQueue struct {
	data       []PQData
	comparator func(PQData, PQData) bool
}

func (pq PriorityQueue) String() string {
	return fmt.Sprintf("PriorityQueue: %v", pq.data)
}

func (pq PriorityQueue) size() int {
	return len(pq.data)
}

func (pq PriorityQueue) isEmpty() bool {
	return pq.size() == 0
}

func (pq PriorityQueue) peek() (int, bool) {
	if pq.isEmpty() {
		return -1, false
	}

	return pq.data[0].val, true
}

func (pq *PriorityQueue) push(data PQData) {
	newData := append(pq.data, data)

	for idx := len(pq.data); idx > 0; {
		p := parent(idx)

		if !pq.comparator(newData[p], newData[idx]) {
			break
		}

		newData[p], newData[idx] = newData[idx], newData[p]
		idx = p
	}

	pq.data = newData
}

func (pq *PriorityQueue) pop() (int, bool) {
	if pq.isEmpty() {
		return -1, false
	}
	q_size, top := pq.size(), pq.data[0]
	if q_size == 1 {
		pq.data = []PQData{}
	} else {
		q_size--
		pq.data[0] = pq.data[q_size]
		pq.data = pq.data[:q_size]
		for idx := 0; idx < q_size; {
			l, r := left(idx), right(idx)
			next := l

			if r < q_size {
				if pq.comparator(pq.data[l], pq.data[r]) {
					next = r
				}
			} else if l >= q_size {
				break
			}

			if pq.comparator(pq.data[idx], pq.data[next]) {
				pq.data[idx], pq.data[next] = pq.data[next], pq.data[idx]
				idx = next
			} else {
				break
			}
		}
	}

	return top.val, true

}

func PriorityQueueMain() {
	pq := PriorityQueue{[]PQData{}, descending}
	fmt.Println("Initial", pq)

	for _, v := range []PQData{{4, 25}, {6, 40}, {5, 35}, {1, 10}, {7, 50}, {2, 15}, {3, 20}} {
		fmt.Println("Pushing", v, "to PriorityQueue")
		pq.push(v)
		fmt.Println(pq)
	}

	for top, exists := pq.pop(); exists; top, exists = pq.pop() {
		fmt.Println("Got top:", top)
		fmt.Println(pq)
	}
}
