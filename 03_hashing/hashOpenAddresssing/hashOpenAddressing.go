package main

import (
	"fmt"
	"hash/fnv"
	"log"
)

type Entry struct {
	key   string
	value int
}

type Hashtable struct {
	table map[int]Entry
	max   int
	num   int
}

func NewHashtable(numObjects int) *Hashtable {
	return &Hashtable{

		table: make(map[int]Entry, numObjects),
		max:   numObjects,
		num:   0,
	}
}

func (h *Hashtable) get(k string) int {
	hashCode := hash(k) % h.max
	for _, bucket := range h.table {
		if bucket.key == k {
			return bucket.value
		}
		hashCode = (hashCode + 1) % h.max
	}

	return 0
}

func (h *Hashtable) put(k string, v int) {
	hashCode := hash(k) % h.max

	_, ok := h.table[hashCode]
	if !ok {
		h.table[hashCode] = Entry{key: k, value: v}
		h.num++
		return
	}

	for i, entry := range h.table {
		if entry.key == k {
			entry.value = v
			return
		}

		hashCode = i+1
	}

	if h.num >= h.max-1 {
		log.Fatal("table is full")
	}

	h.table[hashCode] = Entry{key: k, value: v}
	h.num++
}

func hash(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func main() {
	table := NewHashtable(3)


	table.put("April", 30)
	table.put("May", 31)
	table.put("September", 30)

	fmt.Println(table)

	fmt.Println(table.get("August"))
	fmt.Println(table.get("September"))
}
