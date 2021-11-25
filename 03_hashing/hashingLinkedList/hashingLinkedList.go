package main

import (
	"fmt"
	"hash/fnv"
)

/*
Separate chaining implementation w/ linked lists.
*/

type LinkedBucket struct {
	key string
	value int
	// next int
}

type Hashtable struct {
	table map[int][]LinkedBucket
	size  int
	entry int
}

func NewHashtable(numObjects int) *Hashtable {
	return &Hashtable{

		table: make(map[int][]LinkedBucket, numObjects),
		size:  numObjects,
		entry: 0,
	}
}

func (h *Hashtable) get(k string) int {
	hashCode := hash(k) % h.size

	table := h.table[hashCode]

	for _, bucket := range table {
		if bucket.key == k {
			return bucket.value
		}
	}

	return 0
}

func (h *Hashtable) put(k string, v int) {
	hashCode := hash(k) % h.size

	if len(h.table[hashCode]) < 1 {
		h.table[hashCode] = make([]LinkedBucket, h.size)
	}

	table := h.table[hashCode]

	for _, bucket := range table {
		if bucket.key == k {
			bucket.value = v
			return
		}
	}

	h.table[hashCode][h.entry] = LinkedBucket{key: k, value: v}
	h.entry++
}

func hash(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func main() {
	table := NewHashtable(1000)

	table.put("April", 30)
	table.put("May", 31)
	table.put("September", 30)

	fmt.Println(table.get("August"))
	fmt.Println(table.get("September"))
}
