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

	for h.table[hashCode].key != "" {
		if h.table[hashCode].key == k {
			h.table[hashCode] = Entry{key: k, value: v}
			return
		}
		hashCode++
	}

	if h.num > h.max-1 {
		log.Fatal("table is full")
	}

	h.table[hashCode] = Entry{key: k, value: v}
	h.num++
}

func (h *Hashtable) remove(k string) {
	hashCode := hash(k) % h.max

	entry, ok := h.table[hashCode]
	if !ok {
		return
	}

	if entry.key == k {
		delete(h.table, hashCode)
		h.num--
		return
	}

	for h.table[hashCode].key != "" {
		if h.table[hashCode].key == k {
			delete(h.table, hashCode)
			h.num--
			return
		}
		hashCode++
	}
}

func hash(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func main() {
	table := NewHashtable(12)

	table.put("January", 31)
	fmt.Println(table)
	table.put("February", 28)
	fmt.Println(table)
	table.put("March", 31)
	fmt.Println(table)
	table.put("April", 30)
	fmt.Println(table)
	table.put("May", 31)
	fmt.Println(table)
	table.put("June", 30)
	fmt.Println(table)
	table.put("July", 31)
	fmt.Println(table)
	table.put("August", 31)
	fmt.Println(table)
	table.put("September", 30)
	fmt.Println(table)
	table.put("October", 31)
	fmt.Println(table)
	table.put("November", 30)
	fmt.Println(table)
	table.put("December", 31)
	fmt.Println(table)

	fmt.Println(table.get("August"))
	fmt.Println(table.get("September"))
	fmt.Println(table.get("October"))
	fmt.Println(table.get("April"))
	fmt.Println(table.get("May"))
	fmt.Println(table)

	table.remove("April")
	table.remove("June")
	table.remove("March")
	fmt.Println(table)
}
