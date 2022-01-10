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
	next  *linkedEntry
}

type Hashtable struct {
	table map[int]*linkedEntry
	max   int
	num   int
}

func NewHashtable(numObjects int) *Hashtable {
	if numObjects < 1 {
		panic("hashtable size must be > 0")
	}
	return &Hashtable{

		table: make(map[int]*linkedEntry, numObjects),
		max:   numObjects,
		num:   0,
	}
}

func (h *Hashtable) get(k string) int {
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

func (h *Hashtable) put(k string, v int) {
	hashCode := hash(k) % h.max

	entry, ok := h.table[hashCode]
	if !ok {
		h.table[hashCode] = &linkedEntry{key: k, value: v, next: nil}
	} else {
		h.table[hashCode] = &linkedEntry{key: k, value: v, next: entry}
	}

	h.num++
}

func (h *Hashtable) remove(k string) {
	hashCode := hash(k) % h.max

	entry, ok := h.table[hashCode]
	if !ok {
		return
	}

	if entry.key == k {
		if entry.next == nil {
			delete(h.table, hashCode)
		} else {
			h.table[hashCode] = entry.next
		}
		h.num--
		return
	}

	for entry.next != nil {
		if entry.next.key == k {
			if entry.next.next != nil {
				h.table[hashCode].next = entry.next.next
			} else {
				h.table[hashCode].next = nil
			}
			h.num--
			return
		}
		entry = entry.next
	}
}

func (h *Hashtable) iterate() {
	for _, entry := range h.table {
		fmt.Printf("%s %d, ", entry.key, entry.value)
		for node := entry.next; node != nil; {
			fmt.Printf("%s %d ", node.key, node.value)
			node = node.next
		}
	}
	fmt.Println()
}

func hash(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func main() {
	table := NewHashtable(-100)
	fmt.Println(table)
	//
	// table.put("January", 31)
	// table.put("February", 28)
	// table.put("March", 31)
	// table.put("April", 30)
	// table.put("May", 31)
	// table.put("June", 30)
	// table.put("July", 31)
	// table.put("August", 31)
	// table.put("September", 30)
	// table.put("October", 31)
	// table.put("November", 30)
	// table.put("December", 31)
	//
	// fmt.Println(table.get("August"))
	// fmt.Println(table.get("September"))
	// fmt.Println(table.get("October"))
	// fmt.Println(table.get("April"))
	// fmt.Println(table.get("May"))
	// fmt.Println(table)
	// table.iterate()
	//
	// table.remove("October")
	// table.remove("May")
	// table.remove("September")
	//
	// fmt.Println(table.get("May"))
	// fmt.Println(table.get("September"))
	// fmt.Println(table.get("October"))
	// fmt.Println(table)
	//
	// table.iterate()
}
