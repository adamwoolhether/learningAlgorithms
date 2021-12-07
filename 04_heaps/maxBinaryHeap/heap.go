package main

import "log"

/*
enqueue() is O(log N)
dequeue() is O(log N)
reporting number of entries is O(1)
*/

type Entry struct {
	value    interface{}
	priority int
}

type PriorityQueue struct {
	size    int
	storage []Entry
	number  int
}

func NewPQ(s int) *PriorityQueue {
	return &PriorityQueue{
		size:    s,
		storage: make([]Entry, s+1),
		number:  0,
	}
}

func (pq *PriorityQueue) less(i, j int) bool {
	return pq.storage[i].priority < pq.storage[j].priority
}

// To make a MIN binary heap, reverse the comparison in the less function:
/*
func (pq *PriorityQueue) less(i, j int) bool {
	return pq.storage[i].priority > pq.storage[j].priority
}
*/

func (pq *PriorityQueue) swap(i, j int) {
	pq.storage[i], pq.storage[j] = pq.storage[j], pq.storage[i]
}

func (pq *PriorityQueue) enqueue(v interface{}, p int) {
	if pq.number == pq.size {
		log.Fatalln("PriorityQueue is full")
	}

	pq.number++
	pq.storage[pq.number] = Entry{
		value:    v,
		priority: p,
	}
	pq.swim(pq.number)
}

func (pq *PriorityQueue) swim(child int) {
	for child > 1 && pq.less(child/2, child) {
		pq.swap(child, child/2)
		child /= 2
	}
}

func (pq *PriorityQueue) dequeue() interface{} {
	if pq.number == 0 {
		log.Fatalln("PriorityQueue is empty")
	}

	max := pq.storage[1]
	pq.storage[1] = pq.storage[pq.number]
	pq.storage[pq.number] = Entry{}
	pq.number-- // MUST reduce this BEFORE calling sink()
	pq.sink(1)

	return max.value
}

func (pq *PriorityQueue) sink(parent int) {
	for 2*parent <= pq.number {
		child := 2 * parent
		if child < pq.number && pq.less(child, child+1) {
			child += 1
		}
		if !pq.less(parent, child) {
			break
		}
		pq.swap(child, parent)
		parent = child
	}
}

func main() {

}
