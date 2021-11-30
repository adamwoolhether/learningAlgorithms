package main

import (
	"fmt"
	"hash/fnv"
	"math"
)

type linkedEntry struct {
	key   string
	value int
	next  *linkedEntry
}

type DynamicHashtable struct {
	table map[int]*linkedEntry
	max   int
	num   int

	loadFactor float64
	threshold  float64
}

func NewHashtable(numObjects int, loadFactor float64) *DynamicHashtable {
	return &DynamicHashtable{

		table: make(map[int]*linkedEntry, numObjects),
		max:   numObjects,
		num:   0,

		loadFactor: loadFactor,
		threshold:  math.Min(float64(numObjects)*loadFactor, float64(numObjects)-1),
	}
}

func (h *DynamicHashtable) get(k string) int {
	hashCode := hash(k) % h.max

	entry, ok := h.table[hashCode]
	if !ok {
		return -1
	}

	if entry.key == k {
		return entry.value
	}

	for entry != nil {
		if entry.key == k {
			return entry.value
		}
		entry = entry.next
	}

	return -1
}

func (h *DynamicHashtable) put(k string, v int) {
	hashCode := hash(k) % h.max

	entry, ok := h.table[hashCode]
	if !ok {
		h.table[hashCode] = &linkedEntry{key: k, value: v, next: nil}
	} else {
		h.table[hashCode] = &linkedEntry{key: k, value: v, next: entry}
	}
	h.num++

	if float64(h.num) >= h.threshold {
		h.resize(2*h.max + 1)
	}
}

func (h *DynamicHashtable) resize(newSize int) {
	temp := NewHashtable(newSize, 0.75)
	for _, n := range h.table {
		temp.put(n.key, n.value)
	}
	h.table = temp.table
	h.max = temp.max
	h.threshold = temp.threshold
}

func hash(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func main() {
	table := NewHashtable(6, 0.75)
	table.put("January", 31)
	table.put("February", 28)
	table.put("March", 31)
	table.put("April", 30)
	fmt.Println(table)

	table.put("May", 31)
	table.put("June", 30)
	table.put("July", 31)
	table.put("August", 31)
	table.put("September", 30)
	table.put("October", 31)
	table.put("November", 30)
	table.put("December", 31)
	fmt.Println(table)
}
