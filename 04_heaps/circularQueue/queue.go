package main

import (
	"fmt"
	"log"
)

/*
Circular queue allows for O(1) performance of both enqueue and dequeue.
The first value in the array isn't always storage[0], so we keep track of both the first and last positions.
*/

type Node struct {
	value interface{}
}

type Queue struct {
	size    int
	storage []Node
	first   int
	last    int
	number  int
}

func NewQueue(size int) *Queue {
	return &Queue{
		size:    size,
		storage: make([]Node, size),
		first:   0,
		last:    0,
		number:  0,
	}
}

func (q *Queue) isEmpty() bool {
	return q.number == 0
}

func (q *Queue) isFull() bool {
	return q.number == q.size
}

func (q *Queue) enqueue(v interface{}) {
	node := Node{
		value: v,
	}

	if q.isFull() {
		log.Fatalln("queue is full")
	}

	q.storage[q.last] = node
	q.number++
	q.last = (q.last + 1) % q.size
}

func (q *Queue) dequeue() interface{} {
	if q.isEmpty() {
		log.Fatalln("queue is empty")
	}

	val := q.storage[q.first+1]
	q.number--
	q.first = (q.first + 1) % q.size

	return val
}

func main() {
	people := NewQueue(10)
	fmt.Println(people)

	fmt.Println(people.isEmpty())

	people.enqueue("michael")
	fmt.Println(people)
	fmt.Println(people.isEmpty())
	people.enqueue("joey")
	people.enqueue("ada")
	people.enqueue("sarah")
	fmt.Println(people)

	fmt.Println(people.dequeue())

	fmt.Println(people)
}
