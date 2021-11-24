package main

import (
	"fmt"
	"hash/fnv"
	"log"
)

/*
The performance of put and get are independent of the number of size.
Each is considered to have a constant runtime performance of O(1)
This implementation does not account for collisions.
*/

type Bucket struct {
	key   string
	value int
}

type Hashtable struct {
	table map[int]Bucket
	size  int
}

func NewHashtable(numObjects int) *Hashtable {
	return &Hashtable{

		table: make(map[int]Bucket, numObjects),
		size:  numObjects,
	}
}

func (h *Hashtable) get(k string) int {
	hashCode := hash(k) % h.size

	return h.table[hashCode].value
}

func (h *Hashtable) put(k string, v int) {
	hashCode := hash(k) % h.size

	entry, ok := h.table[hashCode]
	switch {
	case entry.key == k:
		entry.value = v
	case !ok:
		h.table[hashCode] = Bucket{
			key:   k,
			value: v,
		}
	default:
		log.Fatalf("key collision: %s and %s", k, entry.key)
	}
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
