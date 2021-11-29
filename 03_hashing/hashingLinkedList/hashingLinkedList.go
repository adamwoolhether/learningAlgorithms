package main

import (
	"fmt"
	"hash/fnv"
)

/*
Separate chaining implementation w/ linked lists.
*/

type linkedEntry struct {
	key   string
	value int
}

type Hashtable struct {
	table map[int][]linkedEntry
	max   int
	num   int
}

func NewHashtable(numObjects int) *Hashtable {
	return &Hashtable{

		table: make(map[int][]linkedEntry, numObjects),
		max:   numObjects,
		num: 0,
	}
}

func (h *Hashtable) get(k string) int {
	hashCode := hash(k) % h.max

	table := h.table[hashCode]

	for _, bucket := range table {
		if bucket.key == k {
			return bucket.value
		}
	}

	return 0
}

func (h *Hashtable) put(k string, v int) {
	hashCode := hash(k) % h.max

	fmt.Println(k, hashCode)

	entry, ok := h.table[hashCode]
	if !ok {
		h.table[hashCode] = make([]linkedEntry, 1, h.max)
		h.table[hashCode][0] = linkedEntry{key: k, value: v}
		h.num++
		return
	}

	for _, bucket := range entry {
		if bucket.key == k {
			bucket.value = v
			return
		}
	}

	h.table[hashCode] = append([]linkedEntry{{key: k, value: v}}, h.table[hashCode]...)
	// h.num++
}

func (h *Hashtable) remove(k string) {
	hashCode := hash(k) % h.max

	fmt.Println(k, hashCode)

	entry, ok := h.table[hashCode]
	if !ok {
		fmt.Println("not found")
	}

	if len(entry) <2 {
		delete(h.table, hashCode)
		h.num--
		return
	}

	for i, bucket := range entry {
		if bucket.key == k {
			h.table[hashCode] = append(h.table[hashCode][:i], h.table[hashCode][i+1:]...)
			return
		}
	}
}

func hash(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func main() {
	table := NewHashtable(5)

	table.put("April", 30)
	table.put("May", 31)
	table.put("September", 30)
	table.put("RANDY", 33)
	table.put("ERIC", 30)

	fmt.Println(table.get("August"))
	fmt.Println(table.get("September"))
	fmt.Println(table.get("RANDY"))
	fmt.Println(table)

	table.remove("May")
	fmt.Println(table)
}
