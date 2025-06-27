package main

type IndexPair struct {
	i int
	j int
}

type IndexPairQueue []IndexPair

func (q *IndexPairQueue) Enqueue(p ...IndexPair) {
	*q = append(*q, p...)
}

func (q *IndexPairQueue) Dequeue() (IndexPair, bool) {
	if q.isEmpty() {
		return IndexPair{}, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func (q IndexPairQueue) isEmpty() bool {
	return q.size() == 0
}

func (q IndexPairQueue) size() int {
	return len(q)
}
