package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	first *Node
	last  *Node
}

func (q *Queue) isEmpty() bool {
	return q.first == nil
}

func (q *Queue) enqueue(v interface{}) {
	node := Node{
		value: v,
		next:  nil,
	}

	if q.isEmpty() {
		q.first = &node
		q.last = &node
	} else {
		q.last.next = &node
		q.last = q.last.next
	}
}

func (q *Queue) dequeue() interface{} {
	if q.isEmpty() {
		return nil
	}

	val := q.first.value
	q.first = q.first.next

	return val
}

func main() {
	people := Queue{}
	fmt.Println(people.dequeue())

	fmt.Println(people.isEmpty())

	people.enqueue("michael")
	fmt.Println(people.isEmpty())
	people.enqueue("joey")
	people.enqueue("ada")
	people.enqueue("sarah")

	fmt.Println(people.dequeue())

	fmt.Println(people)
}
