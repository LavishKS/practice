package main

type Queue []int

func (q *Queue) Enqueue(p ...int) {
	*q = append(*q, p...)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.isEmpty() {
		return 0, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func (q Queue) isNotEmpty() bool  {
	return q.size() > 0
}

func (q Queue) isEmpty() bool {
	return q.size() == 0
}

func (q Queue) size() int {
	return len(q)
}

type InformTime struct {
	employeeId int
	informTime int
}

type InformTimeQueue []InformTime

func (q *InformTimeQueue) Enqueue(p ...InformTime) {
	*q = append(*q, p...)
}

func (q *InformTimeQueue) Dequeue() (InformTime, bool) {
	if q.isEmpty() {
		return InformTime{}, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func (q InformTimeQueue) isNotEmpty() bool  {
	return q.size() > 0
}

func (q InformTimeQueue) isEmpty() bool {
	return q.size() == 0
}

func (q InformTimeQueue) size() int {
	return len(q)
}
